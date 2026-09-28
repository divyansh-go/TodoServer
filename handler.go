package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func todo(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	switch r.Method {
	case "GET":
		data, _ := json.Marshal(readfile())
		fmt.Fprintln(w, string(data))

	case "POST":
		var f Field
		var t Temp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&t)

		if err != nil || t.Title == nil || t.Completed == nil {

			fmt.Fprintln(w, "Invalid JSON. Please follow the following format:-\n*Example:- {\"Title\":\"non-empty string\",\"Completed\":bool}\n*Ensure both values are passed.\n*ID is assigned by us.")
			break
		}

		if *t.Title == "" {

			fmt.Fprintln(w, "Invalid Input. Title cannot be empty.")
			break

		} else {

			s := readfile()

			newid := rand.Intn(9000) + 1000

		restart:
			for _, val := range s {

				if *val.ID == newid {

					newid = rand.Intn(9000) + 1000
					goto restart
				}
			}

			f.ID = &newid
			f.Title = t.Title
			f.Completed = t.Completed

			s = append(s, f)

			writefile(s)
		}

	default:
		fmt.Fprintln(w, "Please use a valid method.")

	}

	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
}

func todoid(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	var index int
	id, _ := strconv.Atoi(strings.TrimLeft(r.URL.Path, "todos/"))

	for i, val := range readfile() {

		if *val.ID == id {

			index = i
		}
	}

	switch r.Method {
	case "GET":
		for _, val := range readfile() {

			if *val.ID == id {

				data, _ := json.Marshal(val)
				fmt.Fprintln(w, string(data))
			}
		}

	case "PUT":
		var t Temp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&t)

		if err != nil || (t.Title == nil && t.Completed == nil) {

			fmt.Fprintln(w, "Invalid JSON. Please follow the following format.\nExample:- {\"Title\":\"non-empty string\",\"Completed\":bool}\nPlease send atleast either of the two fields or both.")
			break
		}

		if t.Title != nil {

			if *t.Title == "" {

				fmt.Fprintln(w, "Invalid Input. Title cannot be empty.")
				break
			}
		}

		s := readfile()

		if t.Title != nil {

			*s[index].Title = *t.Title
		}

		if t.Completed != nil {

			*s[index].Completed = *t.Completed
		}

		writefile(s)

	case "DELETE":
		data := readfile()

		data = append(data[:index], data[index+1:]...)

		writefile(data)

	default:
		fmt.Fprintln(w, "Please use a valid method.")

	}

	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
}
