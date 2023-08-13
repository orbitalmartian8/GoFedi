package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from Go backend!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
