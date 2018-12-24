package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set hanlder function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})

	// Serve web server with 8080 port
	http.ListenAndServe(":8080", nil)
}
