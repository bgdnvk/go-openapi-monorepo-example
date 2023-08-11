package main

import (
	"net/http"
)

type Server struct{}

func main() {
	server := &Server{}
	router := setupRoutes(server)
	http.ListenAndServe(":8080", router)
}
