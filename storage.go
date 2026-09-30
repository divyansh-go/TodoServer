package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func readfile(x http.ResponseWriter) ([]Field, int) {

	var slc []Field
	jdata, err1 := os.ReadFile("todos.json")
	if err1 != nil {

		log.Println("Failed to read the file.")
		errencode(x, mssg)
	}

	err2 := json.Unmarshal(jdata, &slc)
	if err2 != nil {

		log.Println("Failed to parse the JSON-encoded data.")
		errencode(x, mssg)
	}

	if err1 != nil || err2 != nil {

		x.WriteHeader(http.StatusInternalServerError)
		return slc, 1
	}

	return slc, 0
}

func writefile(slc []Field, x http.ResponseWriter) int {

	jdata, err1 := json.Marshal(slc)
	if err1 != nil {

		log.Println("Failed to encode the given data.")
		errencode(x, mssg)
	}

	err2 := os.WriteFile("todos.json", jdata, 0666)
	if err2 != nil {

		log.Println("Failed to write the data to the file.")
		errencode(x, mssg)
	}

	if err1 != nil || err2 != nil {

		x.WriteHeader(http.StatusInternalServerError)
		return 1
	}
	return 0
}

func errencode(x http.ResponseWriter, str string) {

	x.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(x).Encode(ErrResponse{Error: str})
	if err != nil {

		log.Println("Failed to encode the provided JSON data.")
		x.Header().Set("Content-Type", "plain/text")
		fmt.Fprintln(x, str)
	}
}
