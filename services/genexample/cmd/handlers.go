package main

import (
	"fmt"
	"net/http"
)

func (s *Server) GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (s *Server) GetBye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bye")
}
