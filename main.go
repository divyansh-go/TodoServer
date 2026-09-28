package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/todos", todo)

	for _, val := range readfile() {

		str := "/todos/" + strconv.Itoa(*val.ID)
		http.HandleFunc(str, todoid)
	}

	log.Println("The server started running.......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
