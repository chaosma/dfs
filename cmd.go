package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	//"math/rand"
	"os"
	//"time"
)

func add(filename string) error {

	//	rand.Seed(time.Now().UnixNano())
	//	i1 := rand.Intn(3)
	//	i2 := rand.Intn(3)
	//	fmt.Println(i1, i2)

	fh, err := os.Open(filename)
	if err != nil {
		return err
	}

	sp := new(splitter)
	sp.r = fh
	sp.size = CHUNKSIZE

	fi, err := fh.Stat()
	if err != nil {
		return err
	}
	sp.remain = uint32(fi.Size())

	db, err := bolt.Open(DBName, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = BuildDAG(sp, db)
	if err != nil {
		return err
	}

	fh.Close()
	return nil
}

func getLinks(hash string) error {

	db, err := bolt.Open(DBName, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	n := GetNode(hash, db, BKName)
	for _, l := range n.Links {
		fmt.Printf("%v\t%v\n", l.Hash, l.Tsize)
	}
	return nil
}
