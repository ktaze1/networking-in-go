package tcp

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func DialContext(t *testing.T) {
	deadline := time.Now().Add(5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var d net.Dialer 
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5*time.Second + time.Millisecond) // sleep longer than deadline
		return nil
	}

	conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
	if err != nil {
		conn.Close()
		t.Fatal("connection didn't time out")
	}

	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error isn't a timeout %v", err)
		}
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; actuak: %v", ctx.Err())
	}
}
