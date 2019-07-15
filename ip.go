package main

import (
	"net"
	"strings"
)

func ipList(showIP6 bool) []string {

	cleanIP := func(ip string) string {
		if index := strings.Index(ip, "/"); index != -1 {
			ip = ip[0:index]
		}
		return ip
	}

	IPs := make([]string, 0, 5)

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil
	}

	var urlIP string
	for _, addr := range addrs {
		if strings.Contains(addr.String(), ":") {
			if !showIP6 {
				continue
			}
			urlIP = "[" + cleanIP(addr.String()) + "]"
		} else {
			urlIP = cleanIP(addr.String())
		}

		IPs = append(IPs, urlIP)

	}

	return IPs
}

func startText(ips []string, port string) (result string) {
	println("Server start on:\n=============")
	for _, ip := range ips {
		result += "http://" + ip + ":" + port + "\n"
	}
	return result
}
