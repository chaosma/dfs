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
	cApp.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a file",
			Action: func(c *cli.Context) error {

				filename := c.Args().First()
				err := add(filename)
				return err
			},
		},
		{
			Name:  "links",
			Usage: "get links of an object",
			Action: func(c *cli.Context) error {
				hash := c.Args().First()
				err := getLinks(hash)
				return err
			},
		},
	}

	err := cApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
