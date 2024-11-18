package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handler functions
	mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server is listening on http://localhost:8080")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Could not listen on %s: %v\n", server.Addr, err)
	}
}
