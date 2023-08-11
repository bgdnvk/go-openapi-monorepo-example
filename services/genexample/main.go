package main

import (
	"log"
	"net/http"

	"github.com/bgdnvk/go-openapi-monorepo-example/services/genexample/gen"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	server := &Server{}

	middlewares := []gen.MiddlewareFunc{
		loggingMiddleware,
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("err handling req %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	wrapper := &gen.ServerInterfaceWrapper{
		Handler:            server,
		HandlerMiddlewares: middlewares,
		ErrorHandlerFunc:   errorHandler,
	}

	r.Get("/hello", wrapper.GetHello)
	r.Get("/bye", wrapper.GetBye)

	http.ListenAndServe(":8080", r)
}
