package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Distribution Uses information from foojay JDK Discovery API.
var Distribution = cli.Command{
	Name:        "distribution",
	Aliases:     []string{"dist", "di"},
	Usage:       "Will use the '../distributions' end point of the Foojay Discovery API",
	Description: "dist description",
	Action:      distribution,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "Define the distribution name for example 'zulu', 'oracle'.",
		},
		&cli.StringFlag{
			Name:  "version",
			Usage: "Returns a list of distributions that support the given version.",
		},
	},
}

const url = "https://api.foojay.io/disco/v1.0/distributions"

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

//distribution Analysis the command line options and creates the appropriate URL from it.
func distribution(ctx *cli.Context) error {
	var checkURL = url

	if ctx.IsSet("name") {
		checkURL = fmt.Sprintf("%s/%s", url, ctx.String("name"))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsName(checkURL)
	} else if ctx.IsSet("version") {
		checkURL = fmt.Sprintf("%s/versions/%s", url, ctx.String("version"))
		fmt.Printf("URL: %s\n", checkURL)
		distributionsVersions(checkURL)
	} else {
		fmt.Printf("URL: %s\n", checkURL)
		distributions(checkURL)
	}

	return nil
}

func getData(checkURL string, v interface{}) {
	response, err := http.Get(checkURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("%s\n", response.Status)
		os.Exit(2)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &v)
}

func distributions(checkURL string) error {

	var distributionsStructure DistributionsStructure
	getData(checkURL, &distributionsStructure)

	for i := 0; i < len(distributionsStructure); i++ {
		distribution := distributionsStructure[i]
		fmt.Printf("Name: %s\n", distribution.Name)
		fmt.Printf("API Parameter: %s\n", distribution.APIParameter)
		fmt.Printf("Number of versions: %d\n", len(distribution.Versions))
		for i := 0; i < len(distribution.Versions); i++ {
			fmt.Println(distribution.Versions[i])
		}
	}

	return nil
}

func distributionsName(checkURL string) error {
	var distributionStructure DistributionStructure
	getData(checkURL, &distributionStructure)

	fmt.Printf("Name: %s\n", distributionStructure.Name)
	fmt.Printf("API Parameter: %s\n", distributionStructure.APIParameter)
	fmt.Printf("Number of versions: %d\n", len(distributionStructure.Versions))

	for i := 0; i < len(distributionStructure.Versions); i++ {
		fmt.Println(distributionStructure.Versions[i])
	}
	return nil
}

func distributionsVersions(checkURL string) error {
	var distributionsStructure DistributionsStructure
	getData(checkURL, &distributionsStructure)

	for i := 0; i < len(distributionsStructure); i++ {
		distribution := distributionsStructure[i]
		fmt.Printf("Name: %s\n", distribution.Name)
		fmt.Printf("API Parameter: %s\n", distribution.APIParameter)
		fmt.Printf("Number of versions: %d\n", len(distribution.Versions))
		for i := 0; i < len(distribution.Versions); i++ {
			fmt.Println(distribution.Versions[i])
		}
	}

	return nil
}
