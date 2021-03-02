package dns

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// ListenAndServe listens for a request to retrieve a website's DNS
// information and serves the information back to the client.
func ListenAndServe(udpConn *net.UDPConn) {
	for {
		tmp := make([]byte, 1024)
		_, readAddr, _ := udpConn.ReadFrom(tmp)
		dnsPacket := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacketLayer := dnsPacket.Layer(layers.LayerTypeDNS).(*layers.DNS)
		serve(udpConn, readAddr, dnsPacketLayer)
	}
}
