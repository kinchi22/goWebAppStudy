package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware type
type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		next(c)

		log.Printf("[%s] %q %v]\n",
			c.Request.Method,
			c.Request.URL.String(),
			time.Now().Sub(t))
	}
}

func recoverHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(c.ResponseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next(c)
	}
}

func parseFormHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		c.Request.ParseForm()
		fmt.Println(c.Request.PostForm)
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				c.Params[k] = v[0]
			}
		}
		next(c)
	}
}

func parseJSONBodyHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		var m map[string]interface{}
		if json.NewDecoder(c.Request.Body).Decode(&m); len(m) > 0 {
			for k, v := range m {
				c.Params[k] = v
			}
		}
		next(c)
	}
}
