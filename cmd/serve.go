package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/alethio/ethmock/server"
	"gopkg.in/urfave/cli.v2"
)

// Serve starts a mockup ethereum rpc server
var Serve = &cli.Command{
	Name:  "serve",
	Usage: "serve starts a mocked ethereum rpc server",
	Flags: []cli.Flag{basePath, httpPort},
	Action: func(c *cli.Context) error {
		log.SetLevel(log.DebugLevel)
		srv, err := server.New(c.Int("http-port"), "./testdata/")
		if err != nil {
			log.Fatal(err)
		}

		return srv.Serve()
	},
}
