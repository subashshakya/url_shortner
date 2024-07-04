package main

import (
	"net/http"
)

func NewRouter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/shorten":
			shortenURLHandler(w, r)
		case r.Method == http.MethodGet && len(r.URL.Path) > 1:
			redirectHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}
