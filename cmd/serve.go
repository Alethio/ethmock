package cmd

import (
	"log"

	"github.com/alethio/ethmock/server"
	"gopkg.in/urfave/cli.v2"
)

// Serve starts a mockup ethereum rpc server
var Serve = &cli.Command{
	Name:  "serve",
	Usage: "serve starts a mocked ethereum rpc server",
	Flags: []cli.Flag{basePath, httpPort},
	Action: func(c *cli.Context) error {

		log.Fatal(server.Serve(c.Int("http-port"), "./testdata/"))

		return nil
	},
}
