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
		errencode(x, str)
	}

	err2 := json.Unmarshal(jdata, &slc)
	if err2 != nil {

		log.Println("Failed to parse the JSON-encoded data.")
		errencode(x, str)
	}

	if err1 != nil || err2 != nil {

		return slc, 1
	}

	return slc, 0
}

func writefile(slc []Field, x http.ResponseWriter) {

	jdata, err1 := json.Marshal(slc)
	if err1 != nil {

		log.Println("Failed to encode the given data.")
		errencode(x, str)
	}

	err2 := os.WriteFile("todos.json", jdata, 0666)
	if err2 != nil {

		log.Println("Failed to write the data to the file.")
		errencode(x, str)
	}

}

func errencode(x http.ResponseWriter, str string) {

	err := json.NewEncoder(x).Encode(ErrResponse{Error: str})
	if err != nil {

		log.Println("Failed to encode the provided JSON data.")
		fmt.Fprintln(x, str)
	}
}
