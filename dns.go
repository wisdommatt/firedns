package main

import (
	"log"
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var domains map[string]string

func main() {
	log.Println("server is running")
	domains = map[string]string{
		"timiun.com": "216.58.196.142",
		"meghee.com": "176.32.103.205",
	}

	// getWebsiteIP("cloudnotte.com")

	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}
	udpConn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer udpConn.Close()
	// Waiting to get request on that port
	for {
		tmp := make([]byte, 1024)
		_, readAddr, _ := udpConn.ReadFrom(tmp)
		dnsPacket := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacketLayer := dnsPacket.Layer(layers.LayerTypeDNS).(*layers.DNS)
		serveDNS(udpConn, readAddr, dnsPacketLayer)
	}
}

func serveDNS(conn *net.UDPConn, addr net.Addr, dns *layers.DNS) {
	var dnsResourceRecord layers.DNSResourceRecord
	var ip net.IP
	requestDomain := strings.ToLower(string(dns.Questions[0].Name))
	ipAddress, exist := domains[requestDomain]
	if !exist {
		log.Println("Domain not found !")
		ip, _, _ = net.ParseCIDR("/24")
	} else {
		ip, _, _ = net.ParseCIDR(ipAddress + "/24")
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
	err := dns.SerializeTo(buffer, options)
	if err != nil {
		panic(err)
	}
	conn.WriteTo(buffer.Bytes(), addr)
}

// getWebsiteIP returns a website's IP address.
func getWebsiteIP(website string) (string, error) {
	ips, err := net.LookupIP(website)
	if err != nil {
		return "", err
	}
	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", nil
}
