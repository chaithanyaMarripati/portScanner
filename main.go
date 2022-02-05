package main

import (
	"fmt"
	"math"
	"portScanner/socket"
	"time"
)

//1) get the url and last port number from the user
//2) check for each socket address if connection is established
func main() {
	var url string
	var lastPort int
	fmt.Println("Enter the url/ip of the remote machine")
	fmt.Scanln(&url)
	fmt.Println("Enter the last portNumber(maximum 65535) to be searched for")
	fmt.Scanln(&lastPort)
	lastPort = int(math.Min(float64(lastPort), 65535))
	startTime := time.Now()
	socket.CheckSockets(lastPort, url)
	endTime := time.Now()
	fmt.Println("time taken", endTime.Sub(startTime))
	fmt.Println("press any key to exit")
	fmt.Scanln(&lastPort)
}
