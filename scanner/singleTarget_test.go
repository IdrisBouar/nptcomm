package scanner

import (
	"fmt"
	"testing"
)


func TestLocalPing(t *testing.T) {
	var bench Scanner

	//Test 1
	_, err := bench.NewScanner("192.168.1.2", "80")
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
	}

	if bench.Scan() {
		fmt.Printf("IP Ping is Alive")
	} else {
		fmt.Printf("IP Ping is Dead")
	}

	
}