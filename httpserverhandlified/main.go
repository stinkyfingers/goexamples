package main

import (
	"fmt"
	"log"
	"net/http"
)

// compile & run: go run main.go
// compile: go build -o httpserver main.go

type ourHandler struct {
	message string
}

// main runs http server
func main() {
	// create a servemux
	servemux := http.NewServeMux()

	// ourHandler implements ServeHTTP (below) satisfying the http.Handler interface
	servemux.Handle("/foo", &ourHandler{message: "I'm a foo"})
	servemux.Handle("/bar", &ourHandler{message: "Thar be bar"})

	// listen on port 8080, use servemux
	err := http.ListenAndServe(":8080", servemux)
	log.Fatal(err)
}

func (h *ourHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log defaults to stdout
	log.Print("Method: ", r.Method)

	// print to ResponseWriter
	fmt.Fprint(w, "Your message is: ", h.message)
}
