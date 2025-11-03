package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

func getPort(p *int) (string, error) {
	if checkPort(*p) {
		msg := fmt.Sprintf("Port %d is already in use.", *p)
		return "", errors.New(msg)
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
