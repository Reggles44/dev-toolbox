package handlers

import (
	"dev-server/cmd/server/responses"
	"encoding/json"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

var BasicAuthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != username || pass != password {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		json.NewEncoder(w).Encode(responses.Response{Success: true, Message: "", Error: ""})
	}
})
