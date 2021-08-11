package az

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

var (
	versionString   string = ""
	versionGitSHA   string = ""
	goVersionString string = ""
	buildTimestamp  string = ""
)

func Version() string {
	return fmt.Sprintf("%s (%v, %v, go %v)", versionString, versionGitSHA, buildTimestamp, goVersionString)
}

func PrintVersion(cmd *cli.Cmd) {
	cmd.Spec = ""
	cmd.Action = func() {
		fmt.Println(Version())
	}
}
