package main

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes the given data as a JSON response with the specified status code.
// It sets the "Content-Type" header to "application/json" before writing the response.
//
// Parameters:
//   - w: The http.ResponseWriter to write the response to.
//   - status: The HTTP status code to set for the response.
//   - data: The data to encode as JSON and include in the response body.
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
