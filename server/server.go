package server

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	path string
}

func Serve(port int, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	s := Server{
		path: path,
	}
	ps := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(ps, s)
}
