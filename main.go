package main

import (
	"os"
	"runtime"
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

type kubePARAMETERS struct {
	kubedir     string
	namespace   string
	context     string
	environment string
	rt          string
	domain      string
}

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
			Value: "15",
			Usage: "Specifies sleep duration",
		},
		cli.StringFlag{
			Name:  "phold",
			Value: "5",
			Usage: "Specifies phold duration",
		},
		cli.StringFlag{
			Name:  "namespace",
			Value: "system",
			Usage: "Kube namespace",
		},
		cli.StringFlag{
			Name:  "environment",
			Value: "autosit",
			Usage: "Specifies environment",
		},
		cli.StringFlag{
			Name:  "context",
			Value: "ipt-context",
			Usage: "Specifies context",
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
			Name:    "deploy",
			Aliases: []string{"dep"},
			Usage:   "Triggers the deployment",
			Action: func(deploycli *cli.Context) error {
				deploy()
				return nil
			},
		},
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "Publishes the deployment",
			Action: func(publishcli *cli.Context) error {
				publish()
				return nil
			},
		},
		{
			Name:    "debug",
			Aliases: []string{"deb"},
			Usage:   "Debugs the deployment",
			Action: func(debugcli *cli.Context) error {
				debug()
				return nil
			},
		},
	}

	// sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
