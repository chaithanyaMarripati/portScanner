package socket

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(url string, portNumber int) bool {
	dialer := net.Dialer{Timeout: 1 * time.Second}
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", url, portNumber))
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
func checkPorts(url string, ports []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, port := range ports {
		if isOpen(url, port) {
			fmt.Println("port is open", port)
		}
	}
}
func CheckSockets(lastPort int, url string) {
	var perThread = 3
	var goThreads = lastPort / perThread
	var wg sync.WaitGroup
	for thread := 0; thread < goThreads; thread++ {
		base := thread * 3
		portsCheck := []int{base, base + 1, base + 2}
		wg.Add(1)
		go checkPorts(url, portsCheck, &wg)
	}
	if lastPort%3 != 0 {
		base := goThreads * 3
		open := []int{}
		for i := 0; i < lastPort%3; i++ {
			open = append(open, base+i)
		}
		wg.Add(1)
		go checkPorts(url, open, &wg)
	}
	wg.Wait()
}
