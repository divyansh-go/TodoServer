package main

import (
	"log"
	"net/http"
	"strconv"
)

var id, index int

func main() {

	http.HandleFunc("/todos", todo)

	for i, val := range readfile() {

		id = val.ID
		index = i

		str := "/todos/" + strconv.Itoa(val.ID)
		http.HandleFunc(str, todoid)
	}

	log.Println("The server started running.......")
	http.ListenAndServe(":8080", nil)
}
