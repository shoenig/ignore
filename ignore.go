// Package ignore is used to ignore common nuisance errors.
package ignore

import (
	"io"
	"io/ioutil"
)

// Close c and ignore the error.
func Close(c io.Closer) {
	_ = c.Close()
}

// Drain rc and Close it, ignoring any error.
func Drain(rc io.ReadCloser) {
	_, _ = io.Copy(ioutil.Discard, rc)
	Close(rc)
}
