package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func publish() (method string) {
	method = "This is a placeholder for the publish function"
	fmt.Println(method)
	return
}

func deploy() {
	binary, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		panic(lookErr)
	}

	timeout := "10s"
	args := []string{"kubectl", "get", "pods", "--namespace=kube-system", "--request-timeout=" + timeout}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
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

func context(inputstring string) (outputstring string) {
	if inputstring == "" {
		context := "kube"
		outputstring = "The current context is " + context
	} else {
		context := inputstring
		outputstring = "The current context is " + context
	}
	return
}

func vertical(inputstring string) (outputstring string) {
	if inputstring == "false" {
		vertical := "false"
		outputstring = "vertical is set to " + vertical
	} else {
		vertical := "true"
		outputstring = "vertical is set to " + vertical
	}
	return
}
