package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	// Setup your application (if necessary)
	app := &application{}

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.pingHandler)

	// Serve the request with our handler.
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Define a struct to unmarshal the response JSON.
	var response Response

	// Unmarshal the response body into the response struct.
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Check the response struct values.
	if response.Status != "OK" {
		t.Errorf("Handler returned unexpected status: got %v want OK", response.Status)
	}
}
