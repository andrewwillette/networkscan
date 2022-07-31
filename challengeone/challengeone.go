package challengeone

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"sync"
)

type ActiveIps struct {
	activeIps []string
	mut       sync.Mutex
}

func (aips *ActiveIps) addActiveIp(ip string) {
	aips.mut.Lock()
	defer aips.mut.Unlock()
	aips.activeIps = append(aips.activeIps, ip)
}

const subnetMask = "/24"

// challengeOne Scan local network for all devices.
func challengeOne() []string {
	lip := getLocalIp()

	ip, ipnet, err := net.ParseCIDR(fmt.Sprintf("%s%s", lip, subnetMask))

	if err != nil {
		log.Fatal(err)
	}

	var aips ActiveIps
	var wg sync.WaitGroup
	const max = 20
	semaphore := make(chan struct{}, max)
	checkIp := func(ip string) {
		defer wg.Done()
		if pingIp(ip) {
			aips.addActiveIp(ip)
		}
		<-semaphore
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		semaphore <- struct{}{}
		wg.Add(1)
		go checkIp(ip.String())
	}
	wg.Wait()
	return aips.activeIps
}

// pingIp check if ip responds to ping
func pingIp(ip string) bool {
	out, _ := exec.Command("ping", ip, "-c 1", "-W .05").Output()
	if strings.Contains(string(out), "1 packets received") {
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
