package dnsrecon

import "net"

var hosts []string

func Dnsrecon(targets []string) ([]string, error) {
	for _, target := range targets {
		addrs, err := net.LookupHost(target)
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			hosts = append(hosts, addr)
		}
	}

	return hosts, nil
}
