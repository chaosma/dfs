package main

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

// Data and PBNode are two protobuf message struct
func marshal(by []byte, typ Data_DataType) []byte {
	pbdata := new(Data)
	pbdata.Type = &typ
	pbdata.Data = by
	pbdata.Filesize = proto.Uint64(uint64(len(by))) // return a pointer of integer
	//pbdata.Blocksizes = []uint64{} // blocksizes of child nodes
	enc, err := proto.Marshal(pbdata)
	if err != nil {
		log.Fatal(err)
	}
	pbnode := new(PBNode)
	pbnode.Data = enc
	mdata, err := proto.Marshal(pbnode)
	if err != nil {
		log.Fatal(err)
	}

	//	links := make([]*PBLink, 0, 10)
	//	pbnode.Links = links
	return mdata
}

func marshalSmallFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	mdata := marshal(buf, Data_File)
	mh := multiHash(mdata)

	return mh
}
