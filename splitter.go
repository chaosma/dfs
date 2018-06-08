package main

import (
	"io"
)

var CHUNKSIZE uint32 = 266144 // 256kb

type splitter struct {
	r      io.Reader
	size   uint32 // chunksize
	err    error
	remain uint32 // remaining number of bytes
}

//TODO: add sync.pool buffer to reduce allocation time
func (sp *splitter) NextBytes() ([]byte, error) {
	buf := make([]byte, sp.size)
	n, err := io.ReadFull(sp.r, buf)
	switch err {
	case io.ErrUnexpectedEOF:
		small := make([]byte, n)
		copy(small, buf)
		sp.remain = sp.remain - uint32(n)
		return small, nil
	case nil:
		sp.remain = sp.remain - sp.size
		return buf, nil
	case io.EOF:
		return nil, nil
	default:
		return nil, err
	}
}

func (sp *splitter) done() bool {
	return sp.remain <= 0
}
