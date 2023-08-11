package main

import (
	"log"
	"net/http"

	"github.com/bgdnvk/go-openapi-monorepo-example/services/genexample/api"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(server *Server) *chi.Mux {
	r := chi.NewRouter()

	middlewares := []api.MiddlewareFunc{
		loggingMiddleware,
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("err handling req %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	wrapper := &api.ServerInterfaceWrapper{
		Handler:            server,
		HandlerMiddlewares: middlewares,
		ErrorHandlerFunc:   errorHandler,
	}

	r.Get("/hello", wrapper.GetHello)
	r.Get("/bye", wrapper.GetBye)

	return r
}
