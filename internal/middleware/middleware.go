package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

func ApplyMiddleware(handler http.Handler, middlewareChain ...middleware) http.Handler {
	for _, currentMiddleware := range middlewareChain {
		handler = currentMiddleware(handler)
	}
	return handler
}
