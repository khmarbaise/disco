// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/khmarbaise/disco/modules/helper"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

//Distributions Uses information from foojay JDK Discovery API.
var Distributions = cli.Command{
	Name:        "distributions",
	Aliases:     []string{"dist", "di"},
	Usage:       "Will use the '../distributions' end point of the Foojay Discovery API",
	Description: "dist description",
	Action:      actionDistributions,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  optionName,
			Usage: "Define the distribution name for example 'zulu', 'oracle'.",
		},
		&cli.StringFlag{
			Name:  optionVersion,
			Usage: "Returns a list of distributions that support the given version.",
		},
		&cli.BoolFlag{
			Name:  optionVerbose,
			Usage: "Printout all versions.",
		},
	},
}

//distributionStructure describes the structures under "../distributions/NAME" location.
type distributionStructure struct {
	Name         string   `json:"name"`
	APIParameter string   `json:"api_parameter"`
	Versions     []string `json:"versions"`
}

//distributionsStructure describes the structures under "../distributions" location.
type distributionsStructure []struct {
	distributionStructure
}

type options struct {
	url     string
	verbose bool
}

//distribution Analysis the command line options and creates the appropriate URL from it.
func actionDistributions(ctx *cli.Context) error {
	var checkURL = fmt.Sprintf("%s/distributions", foojayBaseAPI)

	if ctx.IsSet(optionName) {
		checkURL = fmt.Sprintf("%s/%s", checkURL, ctx.String("name"))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsName(options{checkURL, ctx.Bool("verbose")})
	} else if ctx.IsSet(optionVersion) {
		checkURL = fmt.Sprintf("%s/versions/%s", checkURL, ctx.String(optionVersion))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsVersions(checkURL, ctx.Bool(optionVerbose))
	} else {
		fmt.Printf("URL: %s\n", checkURL)
		distributions(checkURL, ctx.Bool(optionVerbose))
	}

	return nil
}

func distributions(checkURL string, verbose bool) error {

	var distributionsStructure distributionsStructure
	helper.GetData(checkURL, &distributionsStructure)

	table := tablewriter.NewWriter(os.Stdout)
	if verbose {
		table.SetHeader([]string{"Name", "API Parameter", "Versions"})
	} else {
		table.SetHeader([]string{"Name", "API Parameter", "Number of Versions"})
	}

	table.SetAutoWrapText(true)
	table.SetRowLine(false)

	for _, v := range distributionsStructure {
		row := []string{}
		if verbose {
			row = []string{v.Name, v.APIParameter, strings.Join(v.Versions, ", ")}
		} else {
			row = []string{v.Name, v.APIParameter, fmt.Sprintf("%d", len(v.Versions))}
		}
		table.Append(row)
	}
	table.Render() // Send output

	return nil
}

func distributionsName(option options) error {
	var distributionStructure distributionStructure
	helper.GetData(option.url, &distributionStructure)

	fmt.Printf("Name: %s\n", distributionStructure.Name)
	fmt.Printf("API Parameter: %s\n", distributionStructure.APIParameter)
	fmt.Printf("Number of versions: %d\n", len(distributionStructure.Versions))

	if option.verbose {
		for i := 0; i < len(distributionStructure.Versions); i++ {
			fmt.Println(distributionStructure.Versions[i])
		}
	}

	return nil
}

func distributionsVersions(checkURL string, verbose bool) error {
	var distributionsStructure distributionsStructure
	helper.GetData(checkURL, &distributionsStructure)

	table := tablewriter.NewWriter(os.Stdout)
	if verbose {
		table.SetHeader([]string{"Name", "API Parameter", "Versions"})
	} else {
		table.SetHeader([]string{"Name", "API Parameter", "Number of Versions"})
	}

	table.SetAutoWrapText(true)
	table.SetRowLine(false)

	for _, v := range distributionsStructure {
		row := []string{}
		if verbose {
			row = []string{v.Name, v.APIParameter, strings.Join(v.Versions, ", ")}
		} else {
			row = []string{v.Name, v.APIParameter, fmt.Sprintf("%d", len(v.Versions))}
		}
		table.Append(row)
	}
	table.Render() // Send output

	return nil
}
