package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/OutOfBedlam/dnazart/internal/az"
	cli "github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("dnazart", "download artifact files from Azure DevOps Pipelines")

	// sub commands
	app.Command("version", "show version", az.PrintVersion)
	app.Command("download", "download artifacts", Download)
	app.Command("hist history", "list builds history", History)

	app.Run(os.Args)
}

func checkPAT() string {
	azPAT := os.Getenv("AZURE_DEVOPS_EXT_PAT")
	if len(azPAT) == 0 {
		fmt.Println("require 'AZURE_DEVOPS_EXT_PAT' environment variable")
		os.Exit(1)
	}
	return azPAT
}

func History(cmd *cli.Cmd) {
	cmd.Spec = "[-n=<limit>] ORG PRJ"
	var (
		pOrg = cmd.StringArg("ORG", "", "organization")
		pPrj = cmd.StringArg("PRJ", "", "project")
		pN   = cmd.IntOpt("n num", 0, "limit number of output, 0: no limit")
	)

	cmd.Action = func() {
		azPAT := checkPAT()

		bl := az.GetBuildList(*pOrg, *pPrj, azPAT)
		currentDefinition := ""
		for i, v := range bl.Value {
			if *pN > 0 && i >= *pN {
				break
			}

			definition := fmt.Sprintf("[%d] Build Definition: %s Type: %s QueueStatus: %s",
				v.Definition.Id, v.Definition.Name, v.Definition.Type, v.Definition.QueueStatus)
			if definition != currentDefinition {
				fmt.Printf("\n%s\n", definition)
				currentDefinition = definition
				sn := len(v.BuildNumber)
				hdrfmt := fmt.Sprintf("BuildId  BuildNumber %%%ds  Status     FinishTime\n", sn-10)
				fmt.Printf(hdrfmt, "")
			}

			status := "?"
			if v.Status == "completed" {
				if v.Result == "succeeded" {
					status = "\x1b[32m✓ Success\x1b[0m"
				} else if v.Result == "failed" {
					status = "\x1b[31m𐄂 Fail\x1b[0m   "
				} else {
					status = fmt.Sprintf("? %s", v.Result)
				}
			} else {
				status = fmt.Sprintf("⥁ %s", v.Status)
			}
			fmt.Printf("%-4d     %s   %s  %s\n", v.Id, v.BuildNumber, status, v.FinishTime)
		}
	}
}

func Download(cmd *cli.Cmd) {
	cmd.Spec = "[-d] [-f] [-o=<output path>] ORG PRJ [BUILDID]"
	var (
		pOrg     = cmd.StringArg("ORG", "", "organization")
		pPrj     = cmd.StringArg("PRJ", "", "project")
		pBuildId = cmd.StringArg("BUILDID", "latest", "build id")
		pDry     = cmd.BoolOpt("d dry", false, "dry-run, print download url without actual downloading")
		pForce   = cmd.BoolOpt("f force", false, "overwrite if output file exists")
		pOut     = cmd.StringOpt("o out", ".", "output directory")
	)

	cmd.Action = func() {
		azPAT := checkPAT()
		buildId := -1

		if *pBuildId == "latest" {
			bl := az.GetBuildList(*pOrg, *pPrj, azPAT)
			if bl.Count == 0 {
				fmt.Printf("No build history found.\n")
				os.Exit(1)
			}
			buildId = bl.Value[0].Id
		} else {
			var err error
			buildId, err = strconv.Atoi(*pBuildId)
			if err != nil {
				fmt.Printf("Invaild BUILDID: %s, %s\n", *pBuildId, err)
				os.Exit(1)
			}
		}

		al, err := az.GetArtifactList(*pOrg, *pPrj, buildId, azPAT)
		if err != nil {
			fmt.Printf("fail to retreive artifacts: %s\n", err)
			os.Exit(1)
		}

		if *pDry {
			for _, a := range al.Value {
				fmt.Printf("%d %s %s\n", a.Id, a.Name, a.Resource.DownloadUrl)
			}
		} else {
			for _, a := range al.Value {
				dst := fmt.Sprintf("%s/%s.zip", *pOut, a.Name)
				if nfo, err := os.Stat(dst); err == nil && nfo != nil && !*pForce {
					fmt.Printf("Destination file already exists, use '-f' to overwrite file: %s\n", dst)
					os.Exit(1)
				}

				err := a.Resource.Download(dst, azPAT)
				if err != nil {
					fmt.Printf("fail to download: %s\n", err)
					os.Exit(1)
				} else {
					fmt.Printf("Downloaded %s", dst)
				}
			}
		}
	}
}
