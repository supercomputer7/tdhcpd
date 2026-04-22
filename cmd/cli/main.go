package main

import (
	"log"

	"github.com/supercomputer7/tdhcpd/pkg/privilege"
	"github.com/supercomputer7/tdhcpd/pkg/dhcp"
)

func startDHCPServer() {
	iface := dhcp.MustSelectInterface()

	handle := dhcp.MustOpenHandle(iface)
	defer handle.Close()

	server := dhcp.NewDHCPServer(handle, iface)

	server.Run()
}

func main() {
	if !privilege.IsElevated() {
		log.Fatal("Not running as superuser/admin")
	}
	startDHCPServer()
}
