
mkdir mock-restapi
cd mock-restapi


Before we start adding code lets initiliase our project

    go mod init github.com/USERNAME/simple-go-service

Create main.go file. and we will listen on port :8000 for web requests.
```
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

Now run this with the following command.

    go build
    go run main.go

This should print out a URL you can navigate to in your browser and see the classic.
Go to http://localhost:8000/hello

Create a Web App
Installing our First Dependency
We're going to install this by running $ go get followed by $ go install:

    #Install a dependency.
    go get -u github.com/gorilla/mux
    go install github.com/gorilla/mux


Docker build image

	cd ./mock-restapi
	docker build -t mock-restapi .

Docker run container.

	docker run -d --rm --name mock-restapi-1 -p 9090:9090 mock-restapi

Endpoint for testing 

	http://localhost:9090/users