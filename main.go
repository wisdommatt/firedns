package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

var records map[string]string

func main() {
	log.Println("running")
	records = map[string]string{
		"google.com": "216.58.196.142",
		"amazon.com": "176.32.103.205",
	}

	//Listen on UDP Port
	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}
	u, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Println(err)
	}
	defer u.Close()

	// Wait to get request on that port
	for {
		tmp := make([]byte, 1024)
		_, readAddr, _ := u.ReadFrom(tmp)
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS).(*layers.DNS)
		serveDNS(u, readAddr, dnsPacket)
	}
}

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	var dnsAnswer layers.DNSResourceRecord
	var ip net.IP
	ipAddress, ok := records[string(request.Questions[0].Name)]
	if !ok {
		log.Println("No domain found !")
		ip, _, _ = net.ParseCIDR("127.0.0.123" + "/24")
		// Todo: Log no data present for the IP and handle:todo
	} else {
		ip, _, _ = net.ParseCIDR(ipAddress + "/24")
	}

	log.Println("the IP")
	log.Println(ip)

	fmt.Println(request.Questions[0])
	fmt.Println(request.Questions[0].Name)

	request.QR = true
	request.ANCount = 1
	request.AA = true
	request.Answers = append(request.Answers, dnsAnswer)
	request.ResponseCode = layers.DNSResponseCodeNoErr
	request.OpCode = layers.DNSOpCodeNotify

	dnsAnswer.IP = ip
	dnsAnswer.Type = layers.DNSTypeA
	dnsAnswer.Name = []byte(request.Questions[0].Name)
	dnsAnswer.Class = layers.DNSClassIN

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{} // See SerializeOptions for more details.
	err := request.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}
