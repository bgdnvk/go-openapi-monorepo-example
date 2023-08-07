package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := Response{Status: "OK"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	})

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
