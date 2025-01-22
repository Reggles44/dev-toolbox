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
	var response responses.BaseResponse

	user, pass, ok := r.BasicAuth()
	if !ok {
		response = responses.BaseResponse{Success: false, Error: "Missing Basic Auth"}
	} else if user != username {
		response = responses.BaseResponse{Success: false, Error: "Username Incorrect"}
	} else if pass != password {
		response = responses.BaseResponse{Success: false, Error: "Password Incorrect"}
	} else {
		response = responses.BaseResponse{Success: true, Message: "Successful Basic Auth"}
	}

	json.NewEncoder(w).Encode(response)
})
