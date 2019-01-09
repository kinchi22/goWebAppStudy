package main

import (
	"net/http"
	"strings"
)

type router struct {
	// key: http method
	// value: HandlerFunc
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	// Check whether map is existing
	m, ok := r.handlers[method]
	if !ok {
		// if no map, create new map
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}

	// Assign URL pattern and handler function on created map
	m[pattern] = h
}

func (r *router) handler() HandlerFunc {
	return func(c *Context) {
		// Check all handlers by http method
		for pattern, handler := range r.handlers[c.Request.Method] {
			if ok, params := match(pattern, c.Request.URL.Path); ok {
				for k, v := range params {
					c.Params[k] = v
				}
				handler(c)
				return
			}
		}

		// Make NotFound error when it fails to find the handler
		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// Skip if pattern and path are same
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			// if pattern is started with ':', put that URL param into params
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}
	}

	return true, params
}
