package cmd

import (
	"github.com/urfave/cli/v2"
)

//EphemeralIds Describe
var EphemeralIds = cli.Command{
	Name:        "ephemeralids",
	Aliases:     []string{"eids", "ids"},
	Usage:       "ephemeralids .....",
	Description: "ephemeralids ....descritpion",
	Action:      ephemeralIds,
}

func ephemeralIds(ctx *cli.Context) error {
	//check.IfError(err)

	return nil
}
