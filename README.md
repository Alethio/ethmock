# The hows and the whys

Needed a simple and fast RPC server that would mock an ethereum client node.
It should be fast to spin up so you can use it in CI and easy to program it's rerturns so you can quickly add more tests.

## Features

This is a list of intended features. WIP.
- [x] http support
- [ ] websockets support
- [ ] helpers for most common fetch requests
- [ ] scriptable requests/responses (ie send pending txs in some order on the subscription) or control subscriptions through some method
- [ ] use as proxy and save (record) all made rpc calls
- [ ] re-fetch jsons you already have so you can see a diff (in case something changed in parity)

## Using it

You have two options of using this mock rpc server:
- [standalone](#standalone) - runs just as a client node would
- [as a package](#as-a-package) - you start it from your tests

### Standalone

### As a package