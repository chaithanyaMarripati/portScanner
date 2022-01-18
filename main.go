package main

import (
	"fmt"
	"portScanner/socket"
)

func main() {
	//get the url and last port number from the user
	//check for each socket address if connection is established
	var url string
	var lastPort int
	fmt.Println("Enter the url/ip of the remote machine")
	fmt.Scanln(&url)
	fmt.Println("Enter the last portNumber(default 1024) to be searched for")
	fmt.Scanln(&lastPort)
	openPorts := []int{}
	for i := 1; i <= lastPort; i++ {
		unixSocket := socket.UnixSocket{
			Url:        url,
			PortNumber: i,
		}
		fmt.Println("checking for port", i)
		if unixSocket.Check() {
			openPorts = append(openPorts, unixSocket.PortNumber)
		}
	}
	fmt.Println("open ports for the remote")
	fmt.Println(openPorts)
}
