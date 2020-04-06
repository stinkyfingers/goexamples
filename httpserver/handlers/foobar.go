package handlers

import "net/http"

// Foo is an http handler
func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("THIS IS FOO"))
}

// Bar is an http handler
func Bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("THIS IS Bar"))
}
