package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func main() {
	fmt.Print(GetOutboundIP())
	fmt.Print(GetHostname())
}
