package dns

import (
	"net"

	"github.com/google/gopacket/layers"
)

// serve returns a website's DNS info to the client that
// requested for it.
// NOTE: this function is unexported because other package(s)
// can't execute it directly.
func serve(conn *net.UDPConn, addr net.Addr, dns *layers.DNS) {}
