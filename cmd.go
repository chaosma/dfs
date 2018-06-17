package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
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
	n, err := GetNode(hash, db, BKName)
	if err != nil {
		return err
	}
	for _, l := range n.Links {
		fmt.Printf("%v\t%v\n", l.Hash, l.Tsize)
	}
	return nil
}

func cat(hash string) error {
	db, err := bolt.Open(DBName, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	n, err := GetNode(hash, db, BKName)
	if err != nil {
		return err
	}
	by := []byte{}
	var arr *[]byte = &by
	err = dfs(db, n, arr)
	if err != nil {
		return err
	}
	fmt.Printf("%s", *arr)
	return nil
}

func dfs(db *bolt.DB, n *Node, arr *[]byte) error {
	if len(n.Links) == 0 {
		pbdata := new(Data)
		proto.Unmarshal(n.Encdata, pbdata)
		*arr = append(*arr, pbdata.Data...)
		return nil
	}

	for _, l := range n.Links {
		m, err := GetNode(l.Hash, db, BKName)
		if err != nil {
			return err
		}
		err = dfs(db, m, arr)
		if err != nil {
			return err
		}
	}
	return nil
}
