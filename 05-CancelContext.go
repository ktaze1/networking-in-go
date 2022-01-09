package tcp

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)



func CancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	go func() {
		defer func() {sync <- struct{}{}}()

		var d net.Dialer
		d.Control = func(_, _ string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}

		conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
		if err != nil {
			t.Log(err)
			return
		}

		conn.Close()
		t.Error("conn didn't timeout")
	}()


	cancel()
	<-sync

	if ctx.Err() != context.Canceled {
		t.Errorf("expected canceled context, actual %q", ctx.Err())
	}
}
