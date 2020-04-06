package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stinkyfingers/goexamples/httpserver/handlers"
)

// init go modules: go mod init github.com/stinkyfingers/goexamples/httpserver
// compile & run: go run main.go
// compile: go build -o httpserver main.go

var (
	// flag variables
	port = flag.Int("p", 8080, "port to listen; defaults to 8080")
)

// main runs http server
func main() {
	// call flag.Parse() right away to instantiate flags
	flag.Parse()

	// call runServer
	err := runServer(*port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runServer(port int) error {
	log.Print("Running server on port ", port)

	// create a servemux
	servemux := http.NewServeMux()

	// server is a pointer to an instance of http.Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: servemux,
	}

	// our handlers reside in the directory "github.com/stinkyfingers/goexamples/httpserver/handlers"
	// function names (e.g. Foo) are capitalized to export
	// specify imported functions by <package_name>.<exported_func_name>
	servemux.HandleFunc("/foo", handlers.Foo)
	servemux.HandleFunc("/bar", handlers.Bar)

	defer server.Close()
	return server.ListenAndServe()
}
