package cmd

import "gopkg.in/urfave/cli.v2"

var Verbose = &cli.BoolFlag{
	Name:  "verbose, v",
	Value: false,
	Usage: "prints debug messages if set",
}

var basePath = &cli.StringFlag{
	Name:  "base-path",
	Value: "testdata",
	Usage: "base data folder",
}

var httpPort = &cli.IntFlag{
	Name:  "http-port",
	Value: 8545,
	Usage: "HTTP port to listen on",
}

var ethClientURL = &cli.StringFlag{
	Name:  "eth-client-url",
	Value: "https://mainnet.infura.io",
	Usage: "Ethereum client rpc url to use for fetching",
}
