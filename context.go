package nova // import "github.com/novakit/nova"

import (
	"context"
	"fmt"
	"net/http"
)

// KeyType custom key type using in context.WithValue() / Context#Value()
type KeyType string

// Context per-request context
type Context struct {
	// Env environment
	Env Env
	// Handlers the handlers
	Handlers []HandlerFunc
	// ErrorHandler must not be nil
	ErrorHandler ErrorHandlerFunc
	// Req the http request
	// Req.Context() will be used for storing pre-request variables
	Req *http.Request
	// Res the http response writer
	Res http.ResponseWriter

	hIndex int // index of current invoking handler
}

// Value get value from request's context
func (c *Context) Value(key string) interface{} {
	return c.Req.Context().Value(KeyType(key))
}

// Set set value for key to request's context, this will update c.Req with a new value
func (c *Context) Set(key string, val interface{}) {
	c.Req = c.Req.WithContext(context.WithValue(c.Req.Context(), KeyType(key), val))
}

// Next invoke the next HandlerFunc registered in application
func (c *Context) Next() {
	// guard panic
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				c.ErrorHandler(c, err)
			} else {
				c.ErrorHandler(c, fmt.Errorf("panic: %v", r))
			}
		}
	}()

	// reached end of handlers chain
	if c.hIndex >= len(c.Handlers) {
		http.NotFound(c.Res, c.Req)
		return
	}
	// save handler index
	i := c.hIndex
	// increase handler index
	c.hIndex++
	// call handler
	if err := c.Handlers[i](c); err != nil {
		// error handling
		c.ErrorHandler(c, err)
	}
}
