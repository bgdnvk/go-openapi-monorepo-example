package main

import (
	"log"
	"net/http"

	petstore "github.com/bgdnvk/go-openapi-monorepo-example/services/swag/petstore"
)

func main() {
	service := &petsService{
		pets: map[int64]petstore.Pet{},
	}

	srv, err := petstore.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}

}
