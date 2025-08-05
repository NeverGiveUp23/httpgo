package ch03

import (
	"errors"
	"net"
	"testing"
	"time"
)

func TestDialTimeout(t *testing.T) {
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		if err := c.Close(); err != nil {
			return
		}
		t.Fatal("connection did not time out")
	}

	var nErr net.Error
	ok := errors.As(err, &nErr)
	if !ok {
		t.Fatal(err)
	}
	if !nErr.Timeout() {
		t.Fatal("error is not a timeout")
	}
}
