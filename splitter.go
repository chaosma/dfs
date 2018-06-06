package main

import (
	"io"
)

type splitter struct {
	r    io.Reader
	size uint32
	err  error
}

//TODO: add sync.pool buffer to reduce allocation time
func (sp *splitter) NextBytes() ([]byte, error) {
	buf := make([]byte, sp.size)
	n, err := io.ReadFull(sp.r, buf)
	switch err {
	case io.ErrUnexpectedEOF:
		small := make([]byte, n)
		copy(small, buf)
		return small, nil
	case nil:
		return buf, nil
	default:
		return nil, err
	}
}
