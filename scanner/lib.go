package scanner 


import (
	"net"
	"fmt"
)


func extractLocal() (net.IP, error) {
    
	var wallet []net.IP

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
        return nil, err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			wallet = append(wallet, ipNet.IP)
		}
	}

	for _, ip := range wallet {
		if ip.To4()[1] == 168 {
			return ip, nil
		}
	}

    return nil, nil
}