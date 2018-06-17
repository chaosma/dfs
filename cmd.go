package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	//"math/rand"
	"os"
	//"time"
)

func add(filename string) {

	//	rand.Seed(time.Now().UnixNano())
	//	i1 := rand.Intn(3)
	//	i2 := rand.Intn(3)
	//	fmt.Println(i1, i2)

	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

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

	fh.Close()
}

func getLinks(hash string) {

	db, err := bolt.Open(DBName, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	n := GetNode(hash, db, BKName)
	for _, l := range n.Links {
		fmt.Printf("%v\t%v\n", l.Hash, l.Tsize)
	}
}
