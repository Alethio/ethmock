package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alethio/ethmock/client"
	"github.com/alethio/ethmock/types"
	log "github.com/sirupsen/logrus"

	"gopkg.in/urfave/cli.v2"
)

// Fetch connects to a server and downloads requested data
var Fetch = &cli.Command{
	Name:  "fetch",
	Usage: "fetch connects to a server and downloads requested data",
	Flags: []cli.Flag{basePath, ethClientURL},
	Action: func(c *cli.Context) error {

		log.Info("starting fetch")

		var args []interface{}
		a := c.Args().Slice()
		switch len(a) {
		case 0:
			log.Fatal("no method given")
		case 1:
			args = []interface{}{}
		case 2:
			err := json.Unmarshal([]byte(a[1]), &args)
			if err != nil {
				log.Fatalf("invalid params json: %s", err)
			}
		default:
			log.Fatal("invalid number of arguments")
		}

		method := a[0]

		rpc, err := client.New(c.String(ethClientURL.Name))
		if err != nil {
			log.Fatal(err)
		}

		req := types.NewJSONRPCRequest(method, args)

		jsonRequest, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		response, err := rpc.Request(jsonRequest)
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		err = json.Indent(&out, response, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		response = out.Bytes()

		base := c.String(basePath.Name)
		folder := filepath.Join(base, req.Path())
		os.MkdirAll(folder, os.ModePerm)

		// write files
		err = ioutil.WriteFile(filepath.Join(folder, "request.json"), jsonRequest, 0644)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(filepath.Join(folder, "response.json"), response, 0644)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	},
}
