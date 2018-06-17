package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	cApp := cli.NewApp()
	cApp.Name = "simpleIPFS"
	cApp.Version = "0.1a"
	cApp.Flags = []cli.Flag{}
	filename := flag.String("filename", "", "a string")
	hash := flag.String("hash", "", "a string")
	addPtr := flag.Bool("add", false, "a bool")
	getPtr := flag.Bool("links", false, "a bool")

	flag.Parse()
	if *addPtr {
		add(*filename)
	}

	if *getPtr {
		getLinks(*hash)
	}

	err := cApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
