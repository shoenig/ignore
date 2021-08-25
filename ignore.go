// Package ignore is used to ignore unimportant errors.
package ignore // import "gophers.dev/pkgs/ignore"

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

// A Rollbacker implements the Rollback method.
type Rollbacker interface {
	Rollback() error
}

// Rollback r and ignore the error.
func Rollback(r Rollbacker) {
	_ = r.Rollback()
}

// Error will ignore the error
func Error(_ error) {
	// do nothing
}

// Panic will cause a panic if the error is not nil.
//
// Intended to pave over cases where an error "cannot" happen.
// e.g. JSON decoding a string literal - if an error occurs, something is
// seriously broken and the program may as well halt.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
