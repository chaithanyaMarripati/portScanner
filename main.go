package main

import (
	"fmt"
	"portScanner/socket"
	"time"
)

//1) get the url and last port number from the user
//2) check for each socket address if connection is established
func main() {
	var url string
	fmt.Println("Enter the url/ip of the remote machine")
	fmt.Scanln(&url)
	startTime := time.Now()
	socket.CheckSockets(url)
	endTime := time.Now()
	fmt.Println("time taken", endTime.Sub(startTime))
	fmt.Println("press any key to exit")
	fmt.Scanln(&url)
}
