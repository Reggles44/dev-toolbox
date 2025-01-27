package handlers

import (
	"dev-server/cmd/server/responses"
	"encoding/json"
	"net/http"
)

var ErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Response{Success: false, Message: "Error", Error: "Generic Error Response"})
})
