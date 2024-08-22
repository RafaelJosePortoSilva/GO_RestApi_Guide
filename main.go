package main

import (
	"apirest/controllers"
	"apirest/routes"
	"fmt"
	"net/http"
)

func main() {
	// Define routes n handler functions
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start server on :8080
	port := 8080
	fmt.Printf("Server running at //localhost:%d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!!!")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!!")
}
