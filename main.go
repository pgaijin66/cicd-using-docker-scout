package main

import (
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "world"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":9090", nil)
}
