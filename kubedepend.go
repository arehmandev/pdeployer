package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func dependencies() {
	validatejson(templateJSON)
	validatejson(deployJSON)
}

func validatejson(jsonfilename string) {

	templateJSONfile, err := ioutil.ReadFile(jsonfilename)
	if err != nil {
		fmt.Println(templateJSON, "not found - please ensure configs/ipt-envconfig-application contains this file")
		panic(err)
	}

	var templateJSONdata interface{}
	PARSEerr := json.Unmarshal(templateJSONfile, &templateJSONdata)
	if PARSEerr != nil {
		fmt.Println("Error parsing JSON")
		panic(PARSEerr)
	}
}
