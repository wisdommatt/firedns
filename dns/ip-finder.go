package dns

import "net"

// getWebsiteIP returns a website's ipv4 IP address.
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
