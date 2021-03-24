package util

import (
	"encoding/json"
	"net/http"
)

// Response : Allows to respond to the client with a corresponding HTTP Status: 200
func Response(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

// ResponseError : Allows to respond to the client with an error message and HTTP Status (40x, 50x)
func ResponseError(w http.ResponseWriter, httpCode int, message string) {
	w.WriteHeader(httpCode)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
