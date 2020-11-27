package middlewares

import "net/http"

type Middleware struct {
	Name         string
	Handler      func(handler http.Handler) http.Handler
	InsertAfter  []string
	InsertBefore []string
	Requires     []string
}
