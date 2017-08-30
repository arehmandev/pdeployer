package main

import (
	"os"
	"runtime"
	"sort"
	"time"

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
	dependencies()
	pdeploy()
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
		cli.StringFlag{
			Name:  "namespace",
			Value: "system",
			Usage: "Kube namespace - defaults to system",
		},
		cli.StringFlag{
			Name:  "",
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
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
