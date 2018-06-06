package main

import (
	"github.com/boltdb/bolt"
	"io"
)

var dbName string = "./dag.go"
var bk string = "dag"

func BuildDAG(sp *splitter, db *bolt.DB) (*Node, error) {
	var root *Node
	for level := 0; sp.done(); level++ {
		nroot := new(Node)
		if root != nil { // nil if it's the first node.
			if err := nroot.AddChild(root, db); err != nil {
				return nil, err
			}
		}
		// fill it up.
		if err := fillNodeRec(sp, db, nroot, level); err != nil {
			return nil, err
		}
		root = nroot
	}
	return root, nil
}

func fillNodeRec(sp *splitter, db *bolt.DB, root *Node, level int) error {
	if level == 0 {
		data, err := sp.NextBytes()
		if err != nil {
			return err
		}
		*root = CreateNode(data, Data_File)
		FlushNode(root, db, bk)
		return nil
	}

	for len(root.Links) < MaximumLinks && !sp.done() {
	}

	return nil
}
