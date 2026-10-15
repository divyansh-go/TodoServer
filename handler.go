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

		w.Header().Set("Content-Type", "application/json")

		slc, n := readfile(w)
		if n == 1 {

			break
		}

		data, err := json.Marshal(slc)
		if err != nil {

			log.Println("Failed to encode the given JSON data.")
			errencode(w, http.StatusInternalServerError, mssg)
			break
		}
		fmt.Fprintln(w, string(data))

	case "POST":

		if r.Header.Get("Content-Type") != "application/json" {

			log.Println("Client Error: Failed to send valid media-type")
			errencode(w, http.StatusUnsupportedMediaType, "Please send a valid media-type along with appropriate Header")
			break
		}

		var f Field
		var t Temp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&t)

		if err != nil || t.Title == nil || t.Completed == nil {

			log.Println("Client error: Failed to provide the proper JSON data.")
			errencode(w, http.StatusBadRequest, "Invalid JSON. Please follow the following format:-\n*Example:- {\"Title\":\"non-empty string\",\"Completed\":bool}\n*Ensure both values are passed.\n*ID is assigned by us.")
			break
		}

		if *t.Title == "" {

			log.Println("Client error: Failed to provide the proper JSON data.")
			errencode(w, http.StatusBadRequest, "Invalid Input. Title cannot be empty.")

			break

		} else {

			slc, n := readfile(w)
			if n == 1 {

				break
			}

			newid := rand.Intn(9000) + 1000

		restart:
			for _, val := range slc {

				if *val.ID == newid {

					newid = rand.Intn(9000) + 1000
					goto restart
				}
			}

			f.ID = &newid
			f.Title = t.Title
			f.Completed = t.Completed

			slc = append(slc, f)

			num := writefile(slc, w)
			if num == 1 {
				break
			}
		}

		w.WriteHeader(http.StatusCreated)

	default:
		log.Println("Client Error: Invalid Method request")
		errencode(w, http.StatusMethodNotAllowed, "Please use a valid method")

	}

	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
}

func todoid(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	var index int

	id, err := strconv.Atoi(strings.TrimLeft(r.URL.Path, "todos/"))
	if err != nil {

		log.Println("Client Error: Failed to send a valid ID.")
		errencode(w, http.StatusBadRequest, "Requested todo does not exist. Please send a valid todo ID.(Numeric Only)")
		log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
		return
	}

	slc, n := readfile(w)
	if n == 1 {

		log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
		return
	}

	for i, val := range slc {

		if *val.ID == id {

			index = i

			goto exists
		}
	}

	log.Println("Client Error: Failed to send a valid ID.")
	errencode(w, http.StatusNotFound, "Requested todo does not exist. Please send a valid todo ID.")
	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))

	return

exists:

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		slc1, n := readfile(w)
		if n == 1 {

			break
		}

		for _, val := range slc1 {

			if *val.ID == id {

				data, err := json.Marshal(val)
				if err != nil {

					log.Println("Failed to encode the given JSON data.")
					errencode(w, http.StatusInternalServerError, mssg)
					log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
					return
				}

				fmt.Fprintln(w, string(data))
			}
		}

	case "PUT":
		if r.Header.Get("Content-Type") != "application/json" {

			log.Println("Client Error: Failed to send valid media-type")
			errencode(w, http.StatusUnsupportedMediaType, "Please send a valid media-type along with appropriate Header")
			break
		}

		var t Temp

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&t)

		if err != nil || (t.Title == nil && t.Completed == nil) {

			log.Println("Client Error: Failed to provide the proper JSON data.")
			errencode(w, http.StatusBadRequest, "Invalid JSON. Please follow the following format.\nExample:- {\"Title\":\"non-empty string\",\"Completed\":bool}\nPlease send atleast either of the two fields or both.")
			break
		}

		if t.Title != nil {

			if *t.Title == "" {

				log.Println("Client Error: Failed to provide the proper JSON data.")
				errencode(w, http.StatusBadRequest, "Invalid Input. Title cannot be empty.")
				break
			}
		}

		s, n := readfile(w)
		if n == 1 {

			break
		}

		if t.Title != nil {

			*s[index].Title = *t.Title
		}

		if t.Completed != nil {

			*s[index].Completed = *t.Completed
		}

		writefile(s, w)

	case "DELETE":
		data, n := readfile(w)
		if n == 1 {

			break
		}

		data = append(data[:index], data[index+1:]...)

		n = writefile(data, w)
		if n == 1 {

			break
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		log.Println("Client Error: Invalid Method request.")
		errencode(w, http.StatusMethodNotAllowed, "Please use a valid method.")

	}

	log.Println(r.Method, r.URL.Path, " Duration ~ ", time.Since(start))
}

func welcome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to to Todo Server!!")
}
