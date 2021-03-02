package main

import (
	"log"
	"net"

	"github.com/wisdommatt/firedns/dns"
)

func main() {
	log.Println("Firedns is running ......")
	// retrieving the UPD connection address.
	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}
	udpConn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer udpConn.Close()
	// Waiting to get request on that port
	dns.ListenAndServe(udpConn)
}
