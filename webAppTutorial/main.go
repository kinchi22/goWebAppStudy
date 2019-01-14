package main

import (
	"fmt"
)

type User struct {
	Id        string
	AddressId string
}

func main() {
	// Make server
	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "Hello world!")
	})
	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})
	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		u := User{Id: c.Params["id"].(string)}
		c.RenderXML(u)
	})
	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJSON(u)
	})
	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})
	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintf(c.ResponseWriter, "create user %v's address\n", c.Params["user_id"])
	})

	s.Run(":8080")
}
