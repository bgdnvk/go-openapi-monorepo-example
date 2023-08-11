package main

import (
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
	wrapper := &gen.ServerInterfaceWrapper{
		Handler:            server,
		HandlerMiddlewares: middlewares,
	}

	r.Get("/hello", wrapper.GetHello)
	r.Get("/bye", wrapper.GetBye)

	http.ListenAndServe(":8080", r)
}
