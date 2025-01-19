package handlers

import (
	"dev-server/cmd/devserver/responses"
	"encoding/json"
	"net/http"
)

var ErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.ErrorResponse)
})


