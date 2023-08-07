package main

import (
	"net/http"
	"time"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
		}()
		next.ServeHTTP(w, r)
	})
}
