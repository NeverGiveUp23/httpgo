package ch03

import (
	"context"
	"errors"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dl)

	defer cancel()

	var d net.Dialer // DialContext is a method on a Dialer
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		// sleep long enough to reach the context deadline.
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		if err := conn.Close(); err != nil {
			return
		}
		t.Fatal("connection did not time out")
	}
	var nErr net.Error
	ok := errors.As(err, &nErr)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not timeout: %v", err)
		}
	}
	if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Errorf("Expected deadline exceeded; actual : %v", ctx.Err())
	}
}
