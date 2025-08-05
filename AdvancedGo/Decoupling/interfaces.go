package main

type reader interface {
	read(b []byte) (int, error)

	badRead(n int) ([]byte, error) // This is not a good example API program in Go, due to []byte now being callable from the bottom of the stack on the user's machine.
}

type file struct {
	name string
}

// read implements the reader interface for a file
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}
