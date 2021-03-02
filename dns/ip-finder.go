package dns

import "net"

// getWebsiteIP returns a website's ipv4 IP address.
func getWebsiteIP(website string) (net.IP, error) {
	ips, err := net.LookupIP(website)
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 != nil {
			return ipv4, nil
		}
	}
	return nil, nil
}
