package handlers

import (
	"dev-server/cmd/devserver/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


var EchoHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
  
  body, err := io.ReadAll(r.Body)
  if err != nil {
    fmt.Printf("could not read body: %s\n", err)
    json.NewEncoder(w).Encode(responses.ErrorResponse)
    return
  }

  w.Write(body)
})
