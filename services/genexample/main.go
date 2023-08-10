package main

import (
	"fmt"
	"net/http"

	"github.com/bgdnvk/go-openapi-monorepo-example/services/genexample/gen"
	"github.com/go-chi/chi/v5"
)

type Server struct{}

func (s *Server) GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Service A!")
}

func main() {
	r := chi.NewRouter()

	var s Server
	gen.HandlerFromMux(&s, r)

	http.ListenAndServe(":8080", r)
}
