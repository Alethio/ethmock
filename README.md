# The hows and the whys

Needed a simple and fast RPC server that would mock an ethereum client node.
It should be fast to spin up so you can use it in CI and easy to program it's rerturns so you can quickly add more tests.

## Features

This is a list of intended features. WIP.
- [x] http support
- [ ] websockets support
- [ ] helpers for most common fetch requests
- [ ] sample data set that could be used direrctly in tests, plug and play as it were...
- [ ] scriptable requests/responses (ie send pending txs in some order on the subscription) or control subscriptions through some method
- [ ] use as proxy and save (record) all made rpc calls
- [ ] re-fetch jsons you already have so you can see a diff (in case something changed in parity)

## Using it

You have two options of using this mock rpc server:
- [standalone](#standalone) - runs just as a client node would
- [as a package](#as-a-package) - you start it from your tests

### Standalone

Using `ethmock` as a standalone binary is easy. All you need to do is

```sh
$ go get github.com/alethio/ethmock
$ ethmock --help
NAME:
   ethmock

USAGE:
   ethmock [global options] command [command options] [arguments...]

VERSION:
...
```
After installing it, you can start it up as a server with the `serve` command, or you can use it's `fetch` command to download sample responses to use from an ethereum  client node.

#### Examples:

#### `serve`
```sh
$ ethmock serve
```

#### `fetch`
```sh
$ ethmock fetch --base-path testdata/infura eth_chainId 
INFO[0000] starting fetch                               
INFO[0000] wrote request.json(90B), response.json(69B) to testdata/infura/eth_chainId

$ ethmock fetch --eth-client-url https://mainnet.infura.io eth_getBlockByNumber '["0xfa1b4",true]'
INFO[0000] starting fetch                               
INFO[0000] wrote request.json(126B), response.json(2.3K) to testdata/eth_getBlockByNumber/0xfa1b4/true 
```

### As a package
Using `ethmock` as a package is mostly used in your tests, so you have a reference, local, ethereum client with preprogrammed responses you can test against.

Example initialization:
```
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/alethio/ethmock/server"
)

func TestRequest(t *testing.T) {
	srv, err := mock.New(8545, "../testdata/mock")
	assert.Nil(t, err)
	go srv.Serve()
	defer srv.Close()

	p, err := httprpc.New("http://localhost:8545")
	assert.Nil(t, err)

	e, err := New(p)
	assert.Nil(t, err)

	n, err := e.GetBlockNumber()
	assert.Nil(t, err)

	assert.Equal(t, int64(7912466), n)
}
```