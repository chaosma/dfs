package main

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
)

func add(filename string) {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	sp := new(splitter)
	sp.r = fh
	sp.size = CHUNKSIZE

	fi, err := fh.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sp.remain = uint32(fi.Size())

	db, err := bolt.Open(DBName, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = BuildDAG(sp, db)
	if err != nil {
		log.Fatal(err)
	}
}
