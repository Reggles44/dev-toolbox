package handlers

import (
	"dev-server/cmd/server/responses"
	"encoding/json"
	"net/http"
)

var SuccessHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Success: true, Message: "", Error: ""})
})

