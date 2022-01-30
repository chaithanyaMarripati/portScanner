package main

import (
	"fmt"
	"math"
	"portScanner/socket"
)

//1) get the url and last port number from the user
//2) check for each socket address if connection is established
func main() {
	var url string
	var lastPort int
	fmt.Println("Enter the url/ip of the remote machine")
	fmt.Scanln(&url)
	fmt.Println("Enter the last portNumber(maximum 1024) to be searched for")
	fmt.Scanln(&lastPort)
	lastPort = int(math.Min(float64(lastPort), 1024))
	openPorts := make([]int, 0, 10)
	for i := 1; i <= lastPort; i++ {
		unixSocket := socket.UnixSocket{
			Url:        url,
			PortNumber: i,
		}
		fmt.Println("checking for port", i)
		if unixSocket.IsOpen() {
			openPorts = append(openPorts, unixSocket.PortNumber)
		}
	}
	fmt.Println("open ports for the remote")
	fmt.Println(openPorts)
}
