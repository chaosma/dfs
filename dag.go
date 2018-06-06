package main

type Node struct {
	data []byte
	hash string
	size uint32
}

type Link struct {
	name string
	hash string
	size uint32
}

func (n *Node) AddChild(that node) error {
}
