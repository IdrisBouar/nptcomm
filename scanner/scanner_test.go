package scanner

import "testing"
import "fmt"


func TestInterface(test *testing.T) {
	if self, addr := IsInterfaceConnected(); self == false {
		test.Errorf("Ethernet or Wi-Fi interface is not connected")
	} else {
		test.Logf("Interface %s is connected", addr)
	}
}


/*Benchmarking function*/
func TestScanNetwork(test *testing.T) {

    activeHosts := ScanNetwork()

    if len(activeHosts) == 0 {
        test.Errorf("No active hosts found")
    }

    fmt.Println("Active hosts:")
    for _, host := range activeHosts {
        fmt.Println(host)
    }
}