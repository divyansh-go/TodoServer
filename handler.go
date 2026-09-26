package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func todo(w http.ResponseWriter, r *http.Request) {

	meth := r.Method

	switch meth {
	case "GET":
		fmt.Fprintln(w, readfile())

	case "POST":
		var f Field
		var s []Field

		json.NewDecoder(r.Body).Decode(&f)

		s = append(s, f)

		writefile(s)
	}
}

func todoid(w http.ResponseWriter, r *http.Request) {

	meth := r.Method

	switch meth {
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
	}
}
