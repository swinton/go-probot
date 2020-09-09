package probot

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Reset will return a new ReadCloser for the body that can be passed to subsequent handlers
func reset(old io.ReadCloser, b []byte) io.ReadCloser {
	old.Close()
	return ioutil.NopCloser(bytes.NewBuffer(b))
}
