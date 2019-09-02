package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"mash.com/mash"
	"mash.com/transport"
	"os"
)

func main() {
	app := cli.App{}

	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "listen",
			Aliases: []string{"l"},
			Usage: "Listen for incoming messages",
			Action: listen,
		},
		{
			Name: "publish",
			Aliases: []string{"p"},
			Usage: "Push a message to a specified peer",
			Action: publish,
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:	"Verbose",
		},
		cli.StringFlag{
			Name:        "Port",
			Value:       "7381",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func publish(c *cli.Context) error {
	verbose := c.Bool("Verbose, v")

	fmt.Print(verbose)
	return nil
}

func listen(c *cli.Context) error {
	self := mash.Init(&mash.Config{}, []transport.Connector{
		&transport.UDPTransport{},
	})

	self.Connect()

	return nil
}
