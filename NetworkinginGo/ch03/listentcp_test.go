package ch03

import (
	"io"
	"net"
	"testing"
)

// Creating listener on 127.0.0.1 using random port(listentcp_test.go)
func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = listener.Close()
	}()

	t.Logf("Bound to %q", listener.Addr())
}

// Testing the process initiating a TCP connection with a server listening to 127.0.0.1 on a random port
func TestDial(t *testing.T) {
	//Create listener on random port
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// channel to pass data between goroutines
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			go func(c net.Conn) {
				// defer for graceful termination before exit
				defer func() {
					if err := c.Close(); err != nil {
						return
					}
					done <- struct{}{}
				}()
				// receive buffer size sending
				buff := make([]byte, 1024)
				for {
					n, err := c.Read(buff)
					if err != nil {
						if err != io.EOF { // EOF -> closing my side of the connection
							t.Error(err)
						}
						return
					}

					t.Logf("recieved: %q", buff[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	// established connection to the listener -> graceful termination from client side
	if err := conn.Close(); err != nil {
		return
	}
	<-done

	if err := listener.Close(); err != nil {
		return
	}
	<-done
}
