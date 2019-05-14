package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetQuote(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(QuoteHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected1 := `"quote":`
	expected2 := `"author":`
	if !strings.Contains(rr.Body.String(), expected1) || !strings.Contains(rr.Body.String(), expected2) {
		t.Errorf("handler returned unexpected body: got %v want %v and %v",
			rr.Body.String(), expected1, expected2)
	}
}
