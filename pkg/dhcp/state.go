package dhcp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type DHCPServer struct {
	handle *pcap.Handle
	source *gopacket.PacketSource
	ipPool chan net.IP
}

func NewDHCPServer(h *pcap.Handle, iface string) *DHCPServer {
	handle, err := pcap.OpenLive(iface, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	return &DHCPServer{
		handle: handle,
		source: source,
		ipPool: make(chan net.IP, 100),
	}
}

func (s *DHCPServer) Run() {
	for packet := range s.source.Packets() {
		event := ParseDHCP(packet)
		if event == nil {
			continue
		}

		switch event.Type {
		case 1: // DISCOVER packet
			s.handleDiscover(event)
		case 3: // REQUEST packet
			s.handleRequest(event)
		}
	}
}

func (s *DHCPServer) handleDiscover(event *DHCPEvent) {
	// TODO: implement DISCOVER → OFFER
}

func (s *DHCPServer) handleRequest(event *DHCPEvent) {
	// TODO: implement REQUEST → ACK
}
