package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"sync"
)

func main() {

}

type ActiveIps struct {
	activeIps []string
	mut       sync.Mutex
}

func (aips *ActiveIps) addActiveIp(ip string) {
	aips.mut.Lock()
	defer aips.mut.Unlock()
	aips.activeIps = append(aips.activeIps, ip)
}

// challengeOne Scan local network for all devices.
func challengeOne() {
	lip := getLocalIp()
	const subnetMask = "/24"
	ip, ipnet, err := net.ParseCIDR(fmt.Sprintf("%s%s", lip, subnetMask))
	if err != nil {
		log.Fatal(err)
	}
	var aips ActiveIps
	var wg sync.WaitGroup
	var counter int
	checkIp := func(ip string) {
		counter++
		fmt.Printf("checking ip %s\n", ip)
		if pingIp(ip) {
			aips.addActiveIp(ip)
		}
		wg.Done()
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		wg.Add(1)
		go checkIp(ip.String())
	}
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
	fmt.Printf("%+v\n", aips.activeIps)
}

// pingIp check if ip responds to ping
func pingIp(ip string) bool {
	// fmt.Printf("Ping ip %s\n", ip)
	out, _ := exec.Command("ping", ip, "-c 1", "-W 1").Output()
	// println(string(out))
	if strings.Contains(string(out), "1 packets received") {
		// fmt.Printf("IP address active: %s\n", ip)
		return true
	} else {
		return false
	}
}

// inc i dont know what this does
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
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
