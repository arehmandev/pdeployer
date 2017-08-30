package main

import "fmt"

func publish() (method string) {
	method = "This is a placeholder for the publish function"
	fmt.Println(method)
	return
}

func deploy() (method string) {
	method = "This is a placeholder for the deploy function"
	fmt.Println(method)
	return
}

func debug() (message string) {
	message = "pdeploy is currently using these settings"
	kubepruneVALUES := []string{"rt", "domain", "namespace", "context", "environment", "environmentBranchOverride", "globalBranchOverride", "deploy", "exclude"}

	fmt.Println(kubepruneVALUES[1])
	// for index := range kubepruneVALUES;  < kubepruneVALUES; ++ {
	//
	// }
	// kubeprune("deploy")

	fmt.Println(message)
	return
}

func context(inputstring string) {
	if inputstring == "" {
		context := "kube"
		fmt.Printf("The current context is %q \n", context)
	} else {
		context := inputstring
		fmt.Printf("The current context is %q \n", context)
	}
}

func vertical(inputstring string) {
	if inputstring == "false" {
		vertical := "false"
		fmt.Printf("The vertical is set to %q \n", vertical)
	} else {
		vertical := "true"
		fmt.Printf("The vertical is set to %q \n", vertical)
	}
}
