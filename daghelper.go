package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

var MaximumLinks int = 174

type FSNode struct {
	Rawdata    []byte
	Filesize   uint64
	Blocksizes []uint64
	Type       Data_DataType
}

type Node struct {
	Encdata    []byte
	Hash       string
	Size       uint64
	Blocksizes []uint64
	Links      []*Link
}

type Link struct {
	Name  string
	Hash  string
	Tsize uint64 // target size
}

func (n *Node) addChild(that *Node) error {
	if len(n.Links) > MaximumLinks {
		return fmt.Errorf("maximum number of links %v reached", MaximumLinks)
	}

	sz := that.getSize()
	//rs := that.Size //same as that.getRawsize()
	//	fmt.Println("rawsize: ", rs)
	//	fmt.Println("cumsize: ", cs)
	//	fmt.Println("totalsize: ", sz)

	l := Link{Name: "", Hash: that.Hash, Tsize: sz}
	n.Links = append(n.Links, &l)
	n.Size += sz
	n.Blocksizes = append(n.Blocksizes, sz)
	return nil
}

func createFSNode(data []byte, typ Data_DataType, filesize uint64, blocksizes []uint64) *FSNode {
	n := new(FSNode)
	n.Rawdata = data
	n.Type = typ
	n.Filesize = filesize
	n.Blocksizes = blocksizes
	return n
}

func (n *FSNode) marshal() []byte {
	pbdata := new(Data)
	typ := n.Type
	pbdata.Type = &typ
	pbdata.Data = n.Rawdata
	pbdata.Filesize = &n.Filesize
	pbdata.Blocksizes = make([]uint64, len(n.Blocksizes))
	copy(pbdata.Blocksizes, n.Blocksizes)
	//pbdata.Blocksizes = append(pbdata.Blocksizes, n.Blocksizes...)
	enc, err := proto.Marshal(pbdata)
	if err != nil {
		log.Fatal(err)
	}
	return enc
}

// leaf node will fill in the data
func createNodeFromFile(data []byte) *Node {
	n := createFSNode(data, Data_File, uint64(len(data)), []uint64{})
	enc := n.marshal()

	pbn := new(Node)
	pbn.Encdata = enc
	pbn.Size = n.Filesize
	pbn.setHash()

	return pbn
}

func (n *Node) getPBNode() *PBNode {
	pbn := &PBNode{}
	if len(n.Links) > 0 {
		pbn.Links = make([]*PBLink, len(n.Links))
	}

	for i, l := range n.Links {
		pbn.Links[i] = &PBLink{}
		pbn.Links[i].Name = &l.Name
		pbn.Links[i].Tsize = &l.Tsize
		if len(l.Hash) != 0 {
			pbn.Links[i].Hash = []byte(l.Hash)
		}
	}

	if len(n.Encdata) > 0 {
		pbn.Data = n.Encdata
	}
	return pbn
}

func (n *Node) getRawsize() uint64 {
	pbdata := new(Data)
	proto.Unmarshal(n.Encdata, pbdata)
	return uint64(len(pbdata.Data))
}

func (n *Node) getCumsize() uint64 {
	s := uint64(len(n.Encdata))
	for _, l := range n.Links {
		s += l.Tsize
	}
	return s
}

func (n *Node) getBlocksize() uint64 {
	s := uint64(len(n.Encdata))
	return s
}

func (n *Node) getSize() uint64 {
	pbn := n.getPBNode()
	by, err := proto.Marshal(pbn)
	if err != nil {
		log.Fatal(err)
	}
	s := uint64(len(by))
	for _, l := range n.Links {
		s += l.Tsize
	}

	return s
}

func (n *Node) calcHash() string {
	pbnode := n.getPBNode()
	mdata, err := proto.Marshal(pbnode)
	if err != nil {
		log.Fatal(err)
	}
	return multiHash(mdata)
}

func (n *Node) setHash() {
	n.Hash = n.calcHash()
}

func (n *Node) getHash() string {
	if len(n.Hash) == 0 {
		n.setHash()
	}
	return n.Hash
}
