package scanner

import (
	"fmt"
	"net"
    "strconv"
)



type Scanner interface {
    NewScanner(string, string) (scannerProfile, error)
    Scan() bool
    InsertInfo(scannerProfile)
}

type scannerProfile struct {

    ip  net.IP
    port    int
    local   bool
}



func (s *scannerProfile) Scan() bool {
    
    switch s.local {
    case true:
        conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: s.ip, Port: s.port})
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            return false
        }
        defer conn.Close()
        return true
    
    case false:
        conn, err := net.Dial("tcp", s.ip.String() + ":" + strconv.Itoa(s.port))
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            return false
        }
        defer conn.Close()
        return true
    }

    return false
}


//Load New Scan operation profile
func (s *scannerProfile)NewScanner(ipInput string, inputPort string) error {
    
    var profile scannerProfile
    var err error

    profile.ip = net.ParseIP(ipInput)
    profile.port, err = strconv.Atoi(inputPort)
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return err
    }
    
    profile.isLocal(extractLocal())

    return nil

}

func (s *scannerProfile) isLocal(compare net.IP, err error) {

    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
    }

    if s.ip.Equal(compare) {
        s.local = true
    } else {
        s.local = false
    }

}