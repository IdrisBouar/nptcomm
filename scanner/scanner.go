package scanner

import (
	"fmt"
	"net"
	"time"
)

// Get the local IP address and subnet mask
func ScanNetwork() []string {
    
    var activeHosts []string
    var localIP net.IP
    var subnetMask net.IPMask


    // Check if Ethernet or Wi-Fi interface is connected

    tmp, localIP := IsInterfaceConnected()
    if !tmp {
        fmt.Println("Ethernet or Wi-Fi interface is not connected")
        panic("Ethernet or Wi-Fi interface is not connected")
    }

    subnetMask = localIP.DefaultMask()

    // Calculate the network address and broadcast address
    network := net.IP(make([]byte, 4))
    for i := 0; i < 4; i++ {
        network[i] = localIP[i] & subnetMask[i]
    }

    broadcast := net.IP(make([]byte, 4))
    for i := 0; i < 4; i++ {
        broadcast[i] = network[i] | ^subnetMask[i]
    }

    // Scan the network
    for ip := network.To4(); ip != nil && ip[3] <= broadcast[3]; incIP(ip) {
        if ip.Equal(localIP) {
            continue
        }

        if hostUp(ip.String()) {
            activeHosts = append(activeHosts, ip.String())
        }
    }

    return activeHosts
}


func incIP(ip net.IP) {
	for j := 3; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func hostUp(ip string) bool {
	conn, err := net.DialTimeout("ip4:icmp", ip, 1000*time.Millisecond)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}

//
func IsInterfaceConnected() (bool, net.IP) {
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println(err)
        return false, nil
    }

    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp != 0 && (iface.Flags&net.FlagLoopback == 0) {
            addrs, err := iface.Addrs()
            if err != nil {
                fmt.Println(err)
                continue
            }

            for _, addr := range addrs {
                if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
                    return true, ipnet.IP
                }
            }
        }
    }

    return false, nil
}



