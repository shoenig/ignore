package ignore

import (
	"errors"
	"io"
	"testing"
)

type fakeCloser struct {
}

var _ io.Closer = (*fakeCloser)(nil)

func (fc *fakeCloser) Close() error {
	return errors.New("some error")
}

func Test_Close(t *testing.T) {
	fc := &fakeCloser{}
	Close(fc)
}

type fakeReadCloser struct {
}

var _ io.ReadCloser = (*fakeReadCloser)(nil)

func (frc *fakeReadCloser) Close() error {
	return errors.New("some error")
}

func (frc *fakeReadCloser) Read([]byte) (int, error) {
	return -1, errors.New("some error")
}

func Test_Drain(t *testing.T) {
	frc := &fakeReadCloser{}
	Drain(frc)
}