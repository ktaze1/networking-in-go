package tcp

import (
	"net"
	"syscall"
	"testing"
	"time"
)

func Timeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d:= net.Dialer { // override
		Control: func(_, addr string, _ syscall.RawConn) error { 
			return &net.DNSError{
				Err: "connection timed out",
				Name: addr,
				Server: "127.0.0.1",
				IsTimeout: true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDial(t *testing.T) {
	c, err := Timeout("tcp", "10.0.0.1:http", 5*time.Second)

	if err != nil {
		c.Close()
		t.Fatal("connection didn't timeout")
	}

	nErr, ok := err.(net.Error)

	if !ok {
		t.Fatal(err)
	}
	if nErr != nil {
		t.Fatal("error is not timeout")
	}
}