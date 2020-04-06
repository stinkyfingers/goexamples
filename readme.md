### simplehttpserver
This is about as simple as they come. It serves on static port 8080 and has paths for `/foo` and `/bar` (no http METHOD). 

`go run main.go`

`curl localhost:8080/foo`

### httpserver 
This server runs on a port that can be specified by flag, creates an `http.Server` object, and imports it's handlers from a package in the `handlers` subdirectory.

`go run main.go -p 6060`

`curl localhost:6060/foo`

### httpserverhandlified
This server illustrates Go's interfaces. The `http` `ServeMux.Handle(pattern string, handler Handler)` takes a `http.Handler` interface as the second parameter. A Handler implements the `ServeHTTP` method (https://golang.org/pkg/net/http/#Handler). In this contrived, simple example, the `ourHandler` implements the `ServeHTTP` method.

`go run main.go`

`curl localhost:8080/foo`

### script
From the `pwd`, lists all dirs w/ LastModified date and files recursively. You may want to avoid running this from root & large directories. 

`go run main.go`

### csvscript
This script reads data from a CSV file into objects and then JSON encodes the data into an output file. In addition to CSV and JSON encoding, it illustrates file I/O (note defer `f.Close()`) and parsing floats from strings. Note the JSON struct tags. 

`go run main.go` 

`cat data.json`

### cli
This tool demonstrates interpreting command line args as well as a simple loop. Go will panic if you reference a slice index that does not exist, so there are length checks in the code. 

`go run main.go 4 foobar`

### clichannels 
This tool does the same as the `cli` above, but uses channels to send a message to a function that is fired off in a goroutine. Note that, without the waitgroup, the program would exit before all messages were printed. 

`go run main.go 4 foobar`

### testexample
This package is a simple what an imported module might look like. The `testexample.go` file contains importable artifacts (note capitalized function and struct names/fields). The `testexample_test.go` file includes a test (note that the filename is suffixed with `_test.go`). 

`go test` - test this package

`go test ./...` - test all recursively

`go test -v -run TestFoobar` - run a specific test, verbose