package io

import (
	"io"
	"os"
)

// OpenFile tries to open a file and returns a reader that is closed automatically on a
// `EOF` error.
func OpenFile(fpath string) (io.Reader, error) {
	f, err := os.Open(fpath)
	return &autoCloser{f}, err
}

// readerCloser is a helper type that automatically closes it's underlying `io.ReaderCloser`
// when reaching an EOF error.
type autoCloser struct {
	io.ReadCloser
}

// Read ...
func (rc *autoCloser) Read(d []byte) (int, error) {
	n, err := rc.ReadCloser.Read(d)

	if err == io.EOF {
		return n, rc.ReadCloser.Close()
	}

	return n, err
}
