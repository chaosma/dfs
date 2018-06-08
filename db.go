package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var DBName string = "./dag.go.db"
var BKName string = "dag"

func encode(n *Node) []byte {
	buf := bytes.Buffer{}
	e := gob.NewEncoder(&buf)
	err := e.Encode(*n)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func decode(b []byte) Node {
	n := Node{}
	buf := bytes.Buffer{}
	buf.Write(b)
	d := gob.NewDecoder(&buf)
	err := d.Decode(&n)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func FlushNode(n *Node, db *bolt.DB, bucketName string) {
	key := []byte(n.Hash)
	value := encode(n)
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		err = bucket.Put(key, value)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

func GetNode(hash string, db *bolt.DB, bucketName string) Node {
	key := []byte(hash)
	var val []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", bucketName)
		}
		val = bucket.Get(key)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return decode(val)
}
