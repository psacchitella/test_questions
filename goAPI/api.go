package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a simple struct for the JSON response
type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// API Controller-like function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Message{
		Status:  "success",
		Message: "Hello from Go API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register routes
	http.HandleFunc("/api/hello", HelloHandler)

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
