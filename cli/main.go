package main

import (
	"github.com/urfave/cli"
	"log"
	"mash.com/mash"
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

	log.Fatal(verbose)
	return nil
}

func listen(c *cli.Context) error {
	peer := mash.NewPeer(mash.Config{})

	peer.Listen()

	return nil
}
