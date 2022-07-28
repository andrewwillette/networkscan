package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

}
func challengeOne() {
	lip := getLocalIp()
	const subnetMask = "/24"
	ip, ipnet, err := net.ParseCIDR(fmt.Sprintf("%s%s", lip, subnetMask))
	if err != nil {
		log.Fatal(err)
	}
	println(ip)
	println(ipnet)
	println(ipnet.Mask)
	println(lip.String())
}

func getLocalIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
