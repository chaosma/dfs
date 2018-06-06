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

func multiHash(by []byte) string {
	h := sha256.New()
	if _, err := io.Copy(h, bytes.NewBuffer(by)); err != nil {
		log.Fatal(err)
	}
	mh := h.Sum(nil)
	header := []byte{0x12, 0x20}
	res := append(header, mh...)
	return base58.Encode(res)
}

func mhMain() {
	filename := os.Args[1]
	by, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	res := multiHash(by)
	fmt.Printf("%s\n", res)

}
