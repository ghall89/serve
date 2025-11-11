package main

import (
	"fmt"
	"net"
	"time"
)

func getPort(p *int) (string, error) {
	if checkPort(*p) {
		port := *p + 1
		msg := fmt.Sprintf("Port %d is already in use. Trying port %d.\r", *p, port)
		fmt.Println(msg)
		return getPort(&port)
	}

	return fmt.Sprintf(":%d", *p), nil
}

func checkPort(port int) bool {
	conn, err := net.DialTimeout("tcp4", fmt.Sprintf("127.0.0.1:%d", port), 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
