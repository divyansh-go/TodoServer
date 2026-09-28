package main

import (
	"encoding/json"
	"os"
)

func readfile() []PtrField {

	var slc []PtrField

	jdata, _ := os.ReadFile("todos.json")
	json.Unmarshal(jdata, &slc)

	return slc
}

func writefile(slc []PtrField) {

	jdata, _ := json.Marshal(slc)
	os.WriteFile("todos.json", jdata, 0666)
}
