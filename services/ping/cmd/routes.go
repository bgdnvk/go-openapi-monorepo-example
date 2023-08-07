package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.pingHandler)

	return app.logRequest(mux)
}
