package main

import (
	"log"
	"net/http"
)

var str = "Failed to perform the required action. Please report the error to the service provider."

func main() {

	http.HandleFunc("/todos", todo)
	http.HandleFunc("/todos/", todoid)

	log.Println("The server started running.......")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
