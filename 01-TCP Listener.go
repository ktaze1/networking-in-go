package tcp

import (
	"net"
	"testing"
)

/*
	Minimal "server"
	should use goroutines in a loop to accept multiple connections
*/
func Listener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0") // accept tcp on 127.0.0.1:0
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}
