package dhcp

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type DHCPEvent struct {
	Type     layers.DHCPMsgType
	Xid      uint32
	ClientIP string
	ClientMAC string
}

func ParseDHCP(packet gopacket.Packet) *DHCPEvent {
	dhcpLayer := packet.Layer(layers.LayerTypeDHCPv4)
	if dhcpLayer == nil {
		return nil
	}

	dhcp := dhcpLayer.(*layers.DHCPv4)

	var msgType layers.DHCPMsgType
	for _, opt := range dhcp.Options {
		if opt.Type == layers.DHCPOptMessageType {
			msgType = layers.DHCPMsgType(opt.Data[0])
		}
	}

	return &DHCPEvent{
		Type:      msgType,
		Xid:       dhcp.Xid,
		ClientMAC: dhcp.ClientHWAddr.String(),
	}
}
