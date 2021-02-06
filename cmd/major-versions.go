package cmd

import (
	"github.com/urfave/cli/v2"
)

//MajorVersions Describe
var MajorVersions = cli.Command{
	Name:        "majorversions",
	Aliases:     []string{"mv"},
	Usage:       "majorversions .....",
	Description: "majorversions ....descritpion",
	Action:      majorVersions,
}

func majorVersions(ctx *cli.Context) error {
	//check.IfError(err)

	return nil
}
