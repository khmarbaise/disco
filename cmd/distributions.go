package cmd

import (
	"fmt"
	"github.com/khmarbaise/disco/modules/helper"
	"github.com/urfave/cli/v2"
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
			Name:  "name",
			Usage: "Define the distribution name for example 'zulu', 'oracle'.",
		},
		&cli.StringFlag{
			Name:  "version",
			Usage: "Returns a list of distributions that support the given version.",
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Printout all versions.",
		},
	},
}

//DistributionStructure describes the structures under "../distributions/NAME" location.
type DistributionStructure struct {
	Name         string   `json:"name"`
	APIParameter string   `json:"api_parameter"`
	Versions     []string `json:"versions"`
}

//DistributionsStructure describes the structures under "../distributions" location.
type DistributionsStructure []struct {
	DistributionStructure
}

type options struct {
	url     string
	verbose bool
}

//distribution Analysis the command line options and creates the appropriate URL from it.
func actionDistributions(ctx *cli.Context) error {
	var checkURL = fmt.Sprintf("%s/distributions", foojayBaseAPI)

	if ctx.IsSet("name") {
		checkURL = fmt.Sprintf("%s/%s", checkURL, ctx.String("name"))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsName(options{checkURL, ctx.Bool("verbose")})
	} else if ctx.IsSet("version") {
		checkURL = fmt.Sprintf("%s/versions/%s", checkURL, ctx.String("version"))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsVersions(checkURL, ctx.Bool("verbose"))
	} else {
		fmt.Printf("URL: %s\n", checkURL)
		distributions(checkURL, ctx.Bool("verbose"))
	}

	return nil
}

func distributions(checkURL string, verbose bool) error {

	var distributionsStructure DistributionsStructure
	helper.GetData(checkURL, &distributionsStructure)

	for i := 0; i < len(distributionsStructure); i++ {
		distribution := distributionsStructure[i]
		fmt.Printf("Name: %16s (API parameter: %16s) Number of versions: %d\n", distribution.Name, distribution.APIParameter, len(distribution.Versions))
		if verbose {
			for version := 0; version < len(distribution.Versions); version++ {
				fmt.Println(distribution.Versions[version])
			}
		}
	}

	return nil
}

func distributionsName(option options) error {
	var distributionStructure DistributionStructure
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
	var distributionsStructure DistributionsStructure
	helper.GetData(checkURL, &distributionsStructure)

	for i := 0; i < len(distributionsStructure); i++ {
		distribution := distributionsStructure[i]
		fmt.Printf("Name: %s\n", distribution.Name)
		fmt.Printf("API Parameter: %s\n", distribution.APIParameter)
		fmt.Printf("Number of versions: %d\n", len(distribution.Versions))
		if verbose {
			for version := 0; version < len(distribution.Versions); version++ {
				fmt.Println(distribution.Versions[version])
			}
		}
	}

	return nil
}
