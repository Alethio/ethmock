package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alethio/ethmock/types"
	log "github.com/sirupsen/logrus"
)

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	log.Debug("post received")
	decoder := json.NewDecoder(r.Body)
	var req types.JSONRPCRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fp := filepath.Join(s.path, req.Path(), "response.json")
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		log.Debugf("path %s does not exist", fp)
		res := fmt.Sprintf(methodNotFoundResponse, req.Method)
		w.Write([]byte(res))
	}

	data, err := ioutil.ReadFile(fp)
	if err != nil {
		// TODO return an error
		log.Error(err)
		return
	}

	var res types.JSONRPCResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		// TODO return an error
		log.Error(err)
		return
	}
	res.ID = req.ID
	ba, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		// TODO return an error
		log.Error(err)
		return
	}
	w.Write(ba)
}
