package dnsrecon

import "net"

func Dnsrecon(targets []string) (map[string][]string, error) {
    resultados := make(map[string][]string)

	for _, target := range targets {
		addrs, err := net.LookupHost(target)
		if err != nil {
			return nil, err
		}

        resultados[target] = addrs
	}

	return resultados, nil
}
