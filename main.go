package main

import (
	"log"
	"net/http"
)

const mssg = "Failed to perform the required action. Please report the error to the service provider."

func main() {

	http.HandleFunc("/todos", todo)
	http.HandleFunc("/todos/", todoid)
	http.HandleFunc("/", welcome)

	log.Println("The server started running.......")
	http.ListenAndServe("localhost:8080", nil)
}
