package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("server listening for requests")
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _:= os.Hostname()
	_, _ = fmt.Fprintf(w, "Hello, world!\n")
	_, _ = fmt.Fprintf(w, "Welcome to K8s training Session\n")
	_, _ = fmt.Fprintf(w, "Hostname: %s\n", host)
}