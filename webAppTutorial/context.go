package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Context type
type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// HandlerFunc will replace http.HandlerFunc
type HandlerFunc func(*Context)

// RenderJSON renders v in JSON format
func (c *Context) RenderJSON(v interface{}) {
	// HTTP Status is StatusOK
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// Content-Type is application/json
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Rendor 'v' in json format
	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// Call RenderErr method if error happened
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

// RenderXML renders v in XML format
func (c *Context) RenderXML(v interface{}) {
	// HTTP Status is StatusOK
	c.ResponseWriter.WriteHeader(http.StatusOK)
	// Content-Type is application/xml
	c.ResponseWriter.Header().Set("Content-Type", "application/xml; charset=utf-8")

	// Rendor 'v' in xml format
	if err := xml.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// Call RenderErr method if error happened
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

// RenderErr renders error if error happens
func (c *Context) RenderErr(code int, err error) {
	if err != nil {
		if code > 0 {
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			// if code is abnormal, set HTTP Status as StatusInternalServerError
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}
