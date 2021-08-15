package main

import (
	"log"
	"net/http"
	"wslight/pkg/server"
)

// Application constants, defining host, port, and protocol.
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	mux := http.NewServeMux()
	//Add handlers
	mux.HandleFunc("/cmd", server.HandleCmd)

	err := http.ListenAndServe(":"+connPort, mux)
	log.Fatal(err)
}
