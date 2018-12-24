package main

import "net/http"

// Context type
type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// HandlerFunc will replace http.HandlerFunc
type HandlerFunc func(*Context)
