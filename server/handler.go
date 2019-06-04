package server

import (
	"encoding/json"
	"net/http"

	"github.com/alethio/ethmock/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var j types.JSONRPCRequest
	err := decoder.Decode(&j)
	if err != nil {
		log.Error(err)
	}
	spew.Dump(j)
}
