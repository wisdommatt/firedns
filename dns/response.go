package dns

import (
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// serve returns a website's DNS info to the client that
// requested for it.
// NOTE: this function is unexported because other package(s)
// can't execute it directly.
// can't execute it directly.
func serve(conn *net.UDPConn, addr net.Addr, dns *layers.DNS) {
	var dnsResourceRecord layers.DNSResourceRecord
	var ip net.IP
	requestDomain := strings.ToLower(string(dns.Questions[0].Name))
	ipAddress, err := getWebsiteIP(requestDomain)
	if err != nil {
		ip, _, _ = net.ParseCIDR("/24")
	} else {
		ip, _, _ = net.ParseCIDR(ipAddress.String() + "/24")
	}
	dnsResourceRecord.IP = ip
	dnsResourceRecord.Name = []byte(requestDomain)
	dnsResourceRecord.Type = layers.DNSTypeA
	dnsResourceRecord.Class = layers.DNSClassIN

	dns.ANCount = 1
	dns.QR = true
	dns.AA = true
	dns.ResponseCode = layers.DNSResponseCodeNoErr
	dns.OpCode = layers.DNSOpCodeNotify
	dns.Answers = append(dns.Answers, dnsResourceRecord)

	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{}
	err = dns.SerializeTo(buffer, options)
	if err != nil {
		panic(err)
	}
	conn.WriteTo(buffer.Bytes(), addr)
}
