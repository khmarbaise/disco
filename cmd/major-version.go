package cmd

import (
	"github.com/urfave/cli/v2"
)

//MajorVersion Describe
var MajorVersion = cli.Command{
	Name:        "majorversion",
	Aliases:     []string{"coi"},
	Usage:       "majorversion .....",
	Description: "majorversion ....descritpion",
	Action:      majorVersion,
}

func majorVersion(ctx *cli.Context) error {
	//check.IfError(err)

	return nil
}
