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

func deploy() (test *kubePARAMETERS) {
	kubestruct := new(kubePARAMETERS)
	kubestruct.kubedir = cwDIR
	kubestruct.namespace = kubeprune("namespace")
	kubestruct.context = kubeprune("context")
	kubestruct.environment = kubeprune("environment")
	kubestruct.rt = kubeprune("rt")
	kubestruct.domain = kubeprune("domain")

	// fmt.Println("Full struct of parse json:", kubestruct)

	binary, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"kubectl", "apply", "-f", cmdAppdir, "--namespace=" + kubestruct.namespace, "--request-timeout=" + cmdPholdflag, "--context=" + kubestruct.context}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	return kubestruct
}

func delete() (test *kubePARAMETERS) {
	kubestruct := new(kubePARAMETERS)
	kubestruct.kubedir = cwDIR
	kubestruct.namespace = kubeprune("namespace")
	kubestruct.context = kubeprune("context")
	kubestruct.environment = kubeprune("environment")
	kubestruct.rt = kubeprune("rt")
	kubestruct.domain = kubeprune("domain")

	// fmt.Println("Full struct of parse json:", kubestruct)

	binary, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"kubectl", "delete", "-f", cmdAppdir, "--namespace=" + kubestruct.namespace, "--request-timeout=" + cmdPholdflag, "--context=" + kubestruct.context}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	return kubestruct
}

func status() (test *kubePARAMETERS) {
	kubestruct := new(kubePARAMETERS)
	kubestruct.kubedir = cwDIR
	kubestruct.namespace = kubeprune("namespace")
	kubestruct.context = kubeprune("context")
	kubestruct.environment = kubeprune("environment")
	kubestruct.rt = kubeprune("rt")
	kubestruct.domain = kubeprune("domain")

	// fmt.Println("Full struct of parse json:", kubestruct)

	binary, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"kubectl", "get", "-f", cmdAppdir, "--namespace=" + kubestruct.namespace, "--request-timeout=" + cmdPholdflag, "--context=" + kubestruct.context}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	return kubestruct
}

func debug() (message string) {
	message = "pdeploy is currently using these settings"
	kubepruneVALUES := []string{"rt", "domain", "namespace", "context", "environment", "environmentBranchOverride", "globalBranchOverride", "deploy", "exclude", cmdAppdir}

	fmt.Println(message)
	for index := 0; index < 9; index++ {
		fmt.Println(kubepruneVALUES[index], "=", kubeprune(kubepruneVALUES[index]))
	}

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
