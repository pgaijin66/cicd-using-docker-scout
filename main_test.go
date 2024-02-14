package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Create a new request
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	handler := http.HandlerFunc(Hello)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "application/json"
	if ct := rr.Header().Get("Content-Type"); ct != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ct, expectedContentType)
	}

	// Check the response body
	expectedBody := map[string]string{"message": "world"}
	var responseBody map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&responseBody); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	if !assert.Equal(t, expectedBody, responseBody) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expectedBody)
	}
}
