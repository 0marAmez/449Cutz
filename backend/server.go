package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, this is a simple Go web server!")
	}

	// Register the handler function for the root route ("/")
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
