package main

import (
	"fmt"
	"log"
	"net/http"
)

// compile & run: go run main.go
// compile: go build -o simplehttpserver main.go

// main runs simple http server
func main() {
	// register patterns to handlers using default servemux
	// http Handlers must have a ServeHTTP method. The http.HandleFunc method wraps the handler for you, such that the Handlers
	// satisfy the http.Handler interface
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)

	// listen on port 8080, use default serve mux (2nd param is nil)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

// handlers
func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FOO"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "BAR")
}
