package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func todo(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, readfile())

	case "POST":
		var f Field
		var s []Field

		json.NewDecoder(r.Body).Decode(&f)

		s = append(s, f)

		writefile(s)
	}

	glb++
	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
}

func todoid(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	switch r.Method {
	case "GET":
		for _, val := range readfile() {

			if val.ID == id {

				fmt.Fprintln(w, val)
			}
		}

	case "PUT":
		var f Field

		json.NewDecoder(r.Body).Decode(&f)

		data := readfile()

		data[index].Title = f.Title
		data[index].Completed = f.Completed

		writefile(data)

	case "DELETE":
		data := readfile()

		data = append(data[:index], data[index+1:]...)

		writefile(data)
	}

	glb++
	log.Println(r.Method, " ", r.URL.Path, " Duration ~ ", time.Since(start))
}
