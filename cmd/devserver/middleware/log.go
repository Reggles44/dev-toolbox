package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQUEST %s\n", r.RemoteAddr)
		fmt.Printf("\t%s %s %s\n", r.Proto, r.Method, r.URL)
		for key, value := range r.Header {
			fmt.Printf("\t%s %s\n", key, strings.Join(value, " "))
		}

    body, err := io.ReadAll(r.Body)
    if err != nil {
      fmt.Printf("Error parsing body: %s", err)
    }
		fmt.Printf("\t%s\n", string(body))
    r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}
