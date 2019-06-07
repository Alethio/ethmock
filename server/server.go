package server

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	path     string
	instance *http.Server
}

func (s *Server) Serve() error {
	return s.instance.ListenAndServe()
}

func (s *Server) Close() error {
	return s.instance.Close()
}

func New(port int, path string) (*Server, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	s := Server{}

	s.path = path

	ps := fmt.Sprintf(":%d", port)
	s.instance = &http.Server{
		Addr:    ps,
		Handler: s,
	}

	return &s, nil
}
