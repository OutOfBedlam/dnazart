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

	app := cli.App("dnazart", "download artifact files from Azure DevOps Pipelines")

	// sub commands
	app.Command("version", "show version", az.PrintVersion)
	app.Command("download", "download artifacts", Download)

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
		bl := az.GetBuildList(*pOrg, *pPrj, azPAT)

		fmt.Printf("Builds count %d\n", bl.Count)
		currentDefinition := ""
		for i, v := range bl.Value {
			definition := fmt.Sprintf("Build Definition[%d] Name:%s Type:%s QueueStatus:%s",
				v.Definition.Id, v.Definition.Name, v.Definition.Type, v.Definition.QueueStatus)
			if definition != currentDefinition {
				fmt.Printf("\n%s\n", definition)
				currentDefinition = definition
			}

			fmt.Printf("%-4d build id:%d status:%s result:%s %s %s\n", i, v.Id, v.Status, v.Result, v.BuildNumber, v.SourceVersion)
		}
	}
}
