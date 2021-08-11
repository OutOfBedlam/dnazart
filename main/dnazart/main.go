package main

import (
	"fmt"
	"os"

	"github.com/OutOfBedlam/dnazart/internal/az"
	cli "github.com/jawher/mow.cli"
)

func main() {
	azPAT := os.Getenv("AZURE_DEVOPS_EXT_PAT")
	if len(azPAT) == 0 {
		fmt.Println("require 'AZURE_DEVOPS_EXT_PAT' environment variable")
		return
	}

	app := cli.App("trkcli", "dtag trkd command line tools")

	// sub commands
	app.Command("version", "show version", az.PrintVersion)
	app.Command("download", "set geofences of facilities", Download)

	app.Run(os.Args)

}

func Download(cmd *cli.Cmd) {
	cmd.Spec = "ORG PRJ"
	var (
		pOrg = cmd.StringArg("ORG", "", "organization")
		pPrj = cmd.StringArg("PRJ", "", "project")
	)
	cmd.Action = func() {
		azPAT := os.Getenv("AZURE_DEVOPS_EXT_PAT")
		az.GetBuildsDefinitions(*pOrg, *pPrj, azPAT)
	}
}
