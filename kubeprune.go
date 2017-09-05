package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func kubeprune(kubestring string) (kubereturn string) {

	templateJSONfile, err := ioutil.ReadFile(deployJSON)
	if err != nil {
		fmt.Println(templateJSON, "not found - please ensure configs/ipt-envconfig-application contains this file")
		panic(err)
	}

	templateJSONdata := string(templateJSONfile)

	switch {
	case kubestring == "rt":
		kubereturn := gjson.Get(templateJSONdata, "rt")
		return kubereturn.String()
	case kubestring == "domain":
		kubereturn := gjson.Get(templateJSONdata, "domain")
		return kubereturn.String()
	case kubestring == "namespace":
		kubereturn := gjson.Get(templateJSONdata, "namespace")
		return kubereturn.String()
	case kubestring == "context":
		kubereturn := gjson.Get(templateJSONdata, "context")
		return kubereturn.String()
	case kubestring == "environment":
		environment := gjson.Get(templateJSONdata, "environment").String()
		return environment
	case kubestring == "environmentBranchOverride":
		environmentBranchOverride := gjson.Get(templateJSONdata, "environmentBranchOverride").String()
		return environmentBranchOverride
	case kubestring == "globalBranchOverride":
		globalBranchOverride := gjson.Get(templateJSONdata, "globalBranchOverride")
		return globalBranchOverride.String()
	case kubestring == "deploy":
		deploy := gjson.Get(templateJSONdata, "deploy.0").String()
		return deploy
	case kubestring == "exclude":
		exclude := gjson.Get(templateJSONdata, "exclude").String()
		return exclude
	}
	return
}
