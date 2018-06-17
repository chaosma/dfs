package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDbSerialization(t *testing.T) {
	filename := "./README.md"
	by, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	mdata := marshal(by, Data_File)
	mh := multiHash(mdata)
	size := uint64(len(by))

	l := Link{Name: "hehe", Hash: "haha", Tsize: 32}

	n := Node{Encdata: mdata, Hash: mh, Size: size, Links: []*Link{&l}}
	by1 := encode(&n)
	n1 := decode(by1)

	if multiHash(n.Encdata) != multiHash(n1.Encdata) {
		t.Fatal("raw data multihash is not the same")
	}
	if n.Links[0].Name != n1.Links[0].Name {
		t.Fatal("link name is not the same")
	}
	fmt.Printf("%s\n%s\n%v\n%s\n%s\n*********\n", multiHash(n.Encdata), n.Hash, n.Size, n.Links[0].Name, n.Links[0].Hash)
	fmt.Printf("%s\n%s\n%v\n%s\n%s\n**********************\n", multiHash(n1.Encdata), n1.Hash, n1.Size, n1.Links[0].Name, n1.Links[0].Hash)

}

func TestDbService(t *testing.T) {
	bucketName := "test"
	dbName := "./dag.db"
	db, err := bolt.Open(dbName, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	l := Link{Name: "hehe", Hash: "haha", Tsize: 32}
	mdata := []byte("this is a test node")
	hash := multiHash(mdata)
	size := uint64(len(mdata))
	n := Node{Encdata: mdata, Hash: hash, Size: size, Links: []*Link{&l}}

	FlushNode(&n, db, bucketName)
	n1, err := GetNode(hash, db, bucketName)
	if err != nil {
		log.Fatal(err)
	}
	if n.Hash != n1.Hash || n.Links[0].Name != n1.Links[0].Name {
		log.Fatal("nodes are not equal")
	}

	fmt.Printf("%s\n%s\n%v\n%s\n%s\n*********\n", multiHash(n.Encdata), n.Hash, n.Size, n.Links[0].Name, n.Links[0].Hash)
	fmt.Printf("%s\n%s\n%v\n%s\n%s\n**********************\n", multiHash(n1.Encdata), n1.Hash, n1.Size, n1.Links[0].Name, n1.Links[0].Hash)

}
