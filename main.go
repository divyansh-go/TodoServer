package main

import (
	"log"
	"net/http"
	"time"
)

const mssg = "Failed to perform the required action. Please report the error to the service provider."

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/todos", todo)
	mux.HandleFunc("/todos/", todoid)
	mux.HandleFunc("/", welcome)

	serverlog := loggingMiddleware(mux)

	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           serverlog,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Println("The server started running.......")
	server.ListenAndServe()
}
