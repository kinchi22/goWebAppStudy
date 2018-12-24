package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})
	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user")
	})
	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user's address")
	})
	r.HandleFunc("POST", "/users/:user_id/addresses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user's address")
	})

	// Serve web server with 8080 port
	http.ListenAndServe(":8080", r)
}
