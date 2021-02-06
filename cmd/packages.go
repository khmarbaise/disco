package cmd

import (
	"github.com/urfave/cli/v2"
)

//Packages ....
var Packages = cli.Command{
	Name:        "packages",
	Aliases:     []string{"pkg"},
	Usage:       "packages on issue",
	Description: "packages ...",
	Action:      packages,
}

func packages(ctx *cli.Context) error {
	///check.IfError(err)

	return nil
}
