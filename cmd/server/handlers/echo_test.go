package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	body := `{"foo": 1, "bar": 2, "foobar": ["hello", "world"]}`

	req, err := http.NewRequest("POST", "/echo/", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != body {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), body)
	}
}
