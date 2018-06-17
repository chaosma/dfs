package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

var bk string = BKName

func BuildDAG(sp *splitter, db *bolt.DB) (*Node, error) {

	var root *Node
	for level := 0; !sp.done(); level++ {
		nroot := new(Node)
		if root != nil { // nil if it's the first node.
			if err := nroot.addChild(root); err != nil {
				return nil, err
			}
		}
		// fill it up.
		if err := fillNodeRec(sp, db, nroot, level); err != nil {
			return nil, err
		}
		root = nroot
	}
	fmt.Printf("%v\t%v\n", root.getHash(), root.getSize())
	return root, nil
}

// make root a balanced tree
func fillNodeRec(sp *splitter, db *bolt.DB, root *Node, level int) error {
	if level == 0 {
		data, err := sp.NextBytes()
		if err != nil {
			return err
		}
		*root = *(createNodeFromFile(data))
		FlushNode(root, db, bk)
		return nil
	}

	for len(root.Links) < MaximumLinks && !sp.done() {
		child := new(Node)
		err := fillNodeRec(sp, db, child, level-1)
		if err != nil {
			return err
		}
		root.addChild(child) // after fill the child node to get its hash full then add
	}
	//createFSNode(data []byte, typ Data_DataType, filesize uint64, blocksizes []uint64)
	n := createFSNode([]byte{}, Data_File, root.Size, root.Blocksizes)
	enc := n.marshal()
	root.Encdata = enc
	root.setHash()
	FlushNode(root, db, bk)

	return nil
}
