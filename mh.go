package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func multiHash(buf []byte) string {
	h := sha256.New()
	if _, err := io.Copy(h, bytes.NewBuffer(buf)); err != nil {
		log.Fatal(err)
	}
	mh := h.Sum(nil)
	header := []byte{0x12, 0x20}
	res := append(header, mh...)
	return base58.Encode(res)
}

func mhMain() {
	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	res := multiHash(buf)
	fmt.Printf("%s\n", res)

}
