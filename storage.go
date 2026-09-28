package main

import (
	"encoding/json"
	"os"
)

func readfile() []Field {

	var slc []Field

	jdata, _ := os.ReadFile("todos.json")
	json.Unmarshal(jdata, &slc)

	return slc
}

func writefile(slc []Field) {

	jdata, _ := json.Marshal(slc)
	os.WriteFile("todos.json", jdata, 0666)
}
