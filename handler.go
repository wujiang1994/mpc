package mpc

import (
	"fmt"
	"net/http"
)

var (
	NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Request(%s %s): not found", r.Method, r.URL.RequestURI()), http.StatusNotFound)
	})

	MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Request(%s %s): method not allowed", r.Method, r.URL.RequestURI()), http.StatusMethodNotAllowed)
	})
)
