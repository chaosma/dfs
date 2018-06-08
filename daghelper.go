package main

import (
	"fmt"
)

var MaximumLinks int = 10

type Node struct {
	Rawdata []byte
	Hash    string
	Size    uint32
	Links   []*Link
}

type Link struct {
	Name string
	Hash string
	Size uint32
}

func (n *Node) AddChild(that *Node) error {
	if len(n.Links) > MaximumLinks {
		return fmt.Errorf("maximum number of links %v reached", MaximumLinks)
	}
	l := Link{Name: "", Hash: that.Hash, Size: that.Size}
	n.Links = append(n.Links, &l)
	return nil
}

func CreateNode(data []byte, typ Data_DataType) Node {
	var n Node
	n.Rawdata = data
	mdata := marshal(data, typ)
	mh := multiHash(mdata)
	n.Hash = mh
	n.Size = uint32(len(data))
	return n
}

func (n *Node) calcHash() {
}
