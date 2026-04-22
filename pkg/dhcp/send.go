package dhcp

import (
	"net"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func parseMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		log.Fatalf("invalid MAC address %q: %v", s, err)
	}
	return mac
}

func (s *DHCPServer) sendOffer(mac string, xid uint32) {
	ip := <-s.ipPool

	eth := layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x02, 0x00, 0x00, 0x00, 0x01, 0x00},
		DstMAC:       parseMAC(mac),
		EthernetType: layers.EthernetTypeIPv4,
	}

	ipLayer := layers.IPv4{
		SrcIP:    net.IPv4(192, 168, 1, 1),
		DstIP:    net.IPv4(255, 255, 255, 255),
		Protocol: layers.IPProtocolUDP,
	}

	udp := layers.UDP{
		SrcPort: 67,
		DstPort: 68,
	}
	udp.SetNetworkLayerForChecksum(&ipLayer)

	dhcp := layers.DHCPv4{
		Operation:    layers.DHCPOpReply,
		Xid:          xid,
		ClientHWAddr: parseMAC(mac),
		YourClientIP:  ip,
	}

	dhcp.Options = append(dhcp.Options,
		layers.NewDHCPOption(layers.DHCPOptMessageType, []byte{byte(layers.DHCPMsgTypeOffer)}),
		layers.NewDHCPOption(layers.DHCPOptServerID, net.IPv4(192,168,1,1).To4()),
		layers.NewDHCPOption(layers.DHCPOptSubnetMask, net.IPv4(255,255,255,0).To4()),
		layers.DHCPOption{Type: layers.DHCPOptEnd},
	)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths: true,
		ComputeChecksums: true,
	}

	gopacket.SerializeLayers(buf, opts,
		&eth, &ipLayer, &udp, &dhcp,
	)

	s.handle.WritePacketData(buf.Bytes())
}
