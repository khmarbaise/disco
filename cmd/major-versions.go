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

//MajorVersions Describe
var MajorVersions = cli.Command{
	Name:        "majorversions",
	Aliases:     []string{"mv"},
	Usage:       "majorversions .....",
	Description: "majorversions ....descritpion",
	Action:      majorVersions,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Major Version  e.g. 1, 5, 9, 11, 17",
		},
		&cli.BoolFlag{
			Name:    "maintained",
			Aliases: []string{"mt"},
			Usage:   "Maintained or not.",
		},
		&cli.BoolFlag{
			Name:    "early-access",
			Aliases: []string{"ea"},
			Usage:   "Early Access.",
		},
		&cli.BoolFlag{
			Name:    "general-availability",
			Aliases: []string{"ga"},
			Usage:   "General availability.",
		},
	},
}

/*
{query} The query to get info about a major version (latest_ea, latest_ga, latest_sts, latest_mts, latest_lts)
*/
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

	var majorVersionsStruct majorVersionsStruct
	helper.GetData(url, &majorVersionsStruct)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Major Version", "Maintained", "Term of Support", "Versions"})
	table.SetAutoWrapText(true)
	table.SetRowLine(true)

	for _, v := range majorVersionsStruct {
		row := []string{fmt.Sprintf("%d", v.MajorVersion), fromBoolToYesNo(v.Maintained), v.TermOfSupport, strings.Join(v.Versions, ", ")}
		table.Append(row)
	}
	table.Render() // Send output
	return nil
}

func fromBoolToYesNo(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}
