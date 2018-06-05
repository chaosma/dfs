package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func marshalMain() {

	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	pbdata := new(Data)
	typ := Data_File
	pbdata.Type = &typ
	pbdata.Data = buf
	pbdata.Filesize = proto.Uint64(uint64(len(buf))) // return a pointer to the integer
	//pbdata.Blocksizes = []uint64{} // blocksizes of child nodes
	fmt.Println("blocksizes:", len(pbdata.Blocksizes))
	fmt.Printf("filesize: %d\n", *pbdata.Filesize)
	enc, err := proto.Marshal(pbdata)
	if err != nil {
		log.Fatal(err)
	}

	pbnode := new(PBNode)
	pbnode.Data = enc
	//	links := make([]*PBLink, 0, 10)
	//	pbnode.Links = links

	mdata, err := proto.Marshal(pbnode)
	if err != nil {
		log.Fatal(err)
	}

	//dd := make(map[string]interface{})
	//dd["data"] = mdata
	//dd["links"] = []string{}
	//js, err := json.Marshal(dd)
	mh := multiHash(mdata)

	fmt.Printf("%s\n", mh)

}
