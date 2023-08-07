package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		app.errorLog.Println("failed to encode response:", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
