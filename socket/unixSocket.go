package socket

import (
	"fmt"
	"net"
	"time"
)

type UnixSocket struct {
	Url        string
	PortNumber int
}

// Check checks if the current socket address is reachable.
// it does this by opening a tcp connection to the unix socket
func (s UnixSocket) Check() bool {
	dialer := net.Dialer{Timeout: 1 * time.Second}
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", s.Url, s.PortNumber))
	if err != nil {
		//tcp connection couldn't be established to the socket
		return false
	}
	//connection is closed and true is returned
	conn.Close()
	return true
}
