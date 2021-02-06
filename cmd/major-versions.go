package cmd

import (
	"fmt"
	"github.com/khmarbaise/disco/modules/helper"
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

//majorVersionsStruct defines the structure which is replied for /major_versions from REST.
type majorVersionsStruct []struct {
	MajorVersion  int      `json:"major_version"`
	TermOfSupport string   `json:"term_of_support"`
	Maintained    bool     `json:"maintained"`
	Versions      []string `json:"versions"`
}

func majorVersions(ctx *cli.Context) error {
	var url = fmt.Sprintf("%s/major_versions", FoojayBaseAPI)

	fmt.Printf("URL: %s\n", url)

	var majorVersionsStrcut majorVersionsStruct
	helper.GetData(url, &majorVersionsStrcut)

	for i := 0; i < len(majorVersionsStrcut); i++ {
		majorVersion := majorVersionsStrcut[i]
		fmt.Printf("Major Version: %d\n", majorVersion.MajorVersion)
		fmt.Printf("Maintained: %v\n", majorVersion.Maintained)
		fmt.Printf("Term of Support: %v\n", majorVersion.TermOfSupport)
	}

	return nil
}
