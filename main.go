package main

import (
	"log"
	"net/http"
)

const mssg = "Failed to perform the required action. Please report the error to the service provider."

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/todos", todo)
	mux.HandleFunc("/todos/", todoid)
	mux.HandleFunc("/", welcome)

	log.Println("The server started running.......")
	http.ListenAndServe("localhost:8080", mux)
}
