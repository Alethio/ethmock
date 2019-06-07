package main

import (
	"os"

	"github.com/alethio/ethmock/cmd"
	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:     "ethmock",
		Usage:    "ethmock command",
		HelpName: "ethmock",
		Version:  "0.0.1",
		Flags: []cli.Flag{
			cmd.Verbose,
		},
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			cmd.Serve,
			cmd.Fetch,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
