package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/error/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ErrorHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"success":false,"message":"Error","error":"Generic Error Response"}`
  body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("handler returned unexpected body:\n\t%v\n\t%v", body, expected)
	}
}
