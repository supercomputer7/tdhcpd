package dhcp

import (
	"log"

	"github.com/google/gopacket/pcap"
)

func MustSelectInterface() string {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range devs {
		log.Println("Interface:", d.Name)
	}

	// naive: pick first real one
	for _, d := range devs {
		if len(d.Addresses) > 0 {
			return d.Name
		}
	}

	log.Fatal("no interface found")
	return ""
}

func MustOpenHandle(iface string) *pcap.Handle {
	handle, err := pcap.OpenLive(
		iface,
		65536,
		true,
		pcap.BlockForever,
	)
	if err != nil {
		log.Fatal(err)
	}

	// DHCP filter
	err = handle.SetBPFFilter("udp and (port 67 or 68)")
	if err != nil {
		log.Fatal(err)
	}

	return handle
}
