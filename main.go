package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

const (
	applicationSRCDIR       = "applications"
	applicationCONFIGDIR    = "configs"
	workspaceDIR            = "generate_workspace"
	configDIR               = workspaceDIR + "/configs"
	applicationDIR          = workspaceDIR + "/applications"
	envCONFIGapp            = "ipt-envconfig-application"
	envCONFIGglobal         = "ipt-envconfig-global"
	componentPROPERTIESfile = "component.override.properties"
	appPROPERTIESfile       = "docker.app.properties"
	templateJSON            = "template.json"
	deployJSON              = "deploy.json"
	cmdAPPLICATIONname      = ""
)

var (
	kubeNAMESPACE       string
	kubeCONTEXT         string
	kubeDOMAIN          string
	kubeENVIRONMENT     string
	typecores           = runtime.NumCPU()
	cmdSLEEPS           = 30
	cmdSLEEPSspec       = true
	optionalDeploy      = ""
	optionalPublish     = ""
	optionalContext     = ""
	optionalDebug       = ""
	cmdVERTICALoverride = false
	cmdDELAY            = 1
	cwDIR, cwDIRerr     = os.Getwd()
)

func main() {
	kubeprune("deploy")
	// dependencies()
	// pdeploy()
}

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
		fmt.Println(kubereturn)
	case kubestring == "domain":
		kubereturn := gjson.Get(templateJSONdata, "domain")
		fmt.Println(kubereturn)
	case kubestring == "namespace":
		kubereturn := gjson.Get(templateJSONdata, "namespace")
		fmt.Println(kubereturn)
	case kubestring == "context":
		kubereturn := gjson.Get(templateJSONdata, "context")
		fmt.Println(kubereturn)
	case kubestring == "environment":
		kubereturn := gjson.Get(templateJSONdata, "environment")
		fmt.Println(kubereturn)
	case kubestring == "environmentBranchOverride":
		kubereturn := gjson.Get(templateJSONdata, "environmentBranchOverride")
		fmt.Println(kubereturn)
	case kubestring == "globalBranchOverride":
		kubereturn := gjson.Get(templateJSONdata, "globalBranchOverride")
		fmt.Println(kubereturn)
	case kubestring == "deploy":
		kubereturn := gjson.Get(templateJSONdata, "deploy.0").Array()
		fmt.Println(kubereturn)
	case kubestring == "exclude":
		kubereturn := gjson.Get(templateJSONdata, "exclude").Array()
		fmt.Println(kubereturn)
	}
	return kubereturn
}

func pdeploy() {
	app := cli.NewApp()
	app.Name = "pdeployer"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Abdul Rehman",
			Email: "rehma017@outlook.com",
		},
	}
	app.Copyright = "(c) 2017 Abdul Rehman"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "sleep",
			Value: "0",
			Usage: "Specifies sleep duration, defaults to 0",
		},
		cli.StringFlag{
			Name:  "phold",
			Value: "0",
			Usage: "Specifies phold duration, defaults to 0",
		},
	}

	// if c.Args().Get(0)) == 0 {
	// 	app.Action = func(c *cli.Context) error {
	// 		sleeptime := c.Args().Get(0)
	// 		fmt.Println("Sleep duration is", sleeptime)
	// 		return
	// 	}
	// } else {
	// 	sleeptime := c.Args().Get(0)
	// 	fmt.Println("Sleep duration is ", sleeptime)
	// 	return
	// }

	app.Commands = []cli.Command{
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "(Optional) - Publishes the deployment",
			Action: func(publishcli *cli.Context) error {
				publish()
				return nil
			},
		},
		{
			Name:    "deploy",
			Aliases: []string{"dep"},
			Usage:   "(Optional) - Triggers the deployment",
			Action: func(deploycli *cli.Context) error {
				deploy()
				return nil
			},
		},
		{
			Name:    "debug",
			Aliases: []string{"deb"},
			Usage:   "(Optional) - Debugs the deployment",
			Action: func(debugcli *cli.Context) error {
				debug()
				return nil
			},
		},
		{
			Name:    "context",
			Aliases: []string{"c"},
			Usage:   "Deploys to the context specified. Defaults to 'kube'",
			Action: func(contextcli *cli.Context) error {
				context(contextcli.Args().Get(0))
				return nil
			},
		},
		{
			Name:    "vertical",
			Aliases: []string{"v"},
			Usage:   "Specifies the vertical. Defaults to 'true'",
			Action: func(verticalcli *cli.Context) error {
				vertical(verticalcli.Args().Get(0))
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}

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

func debug() (method string) {
	method = "This is a placeholder for the debug function"
	fmt.Println(method)
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
