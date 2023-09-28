package main

import (
	"log"
	"net"
	"net/http"

	"github.com/brunoderisso/start-crud-go/greeting"
)

func main() {
	// set a HTTP request handle function for path /greeting and registrate it
	http.HandleFunc("/greeting", greeting.Greeting)

	// print out the server is going to start and show the time
	log.Println("Starting server....")

	// create server at localhost:8080 and using tcp as the network
	listener, err := net.Listen("tcp", ":8080")

	// if recieve error, record it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// setup HTTP connection for the listener of the server
	http.Serve(listener, nil)
}
