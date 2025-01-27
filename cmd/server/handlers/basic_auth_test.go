package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthHandler(t *testing.T) {
	tests := []struct {
		username string
		password string
		status   int
	}{
		{username: "abc", password: "123", status: http.StatusOK},
		{username: "foo", password: "bar", status: http.StatusUnauthorized},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			req, err := http.NewRequest("POST", "/basic_auth/", nil)
			req.SetBasicAuth(test.username, test.password)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(BasicAuthHandler)
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != test.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusNoContent)
			}
		})
	}
}
