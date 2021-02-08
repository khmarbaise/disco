// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/khmarbaise/disco/modules/check"
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
	Usage:       "",
	Description: "Majorversions will access the endpoint ../majorversions of the disco API.",
	Action:      actionMajorVersions,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    optionMajorVersion,
			Aliases: []string{"v"},
			Usage:   "Major Version  e.g. 1, 5, 9, 11, 17",
		},
		&cli.StringFlag{
			Name:    optionLatest,
			Aliases: []string{"l"},
			Usage:   "latest_ea ('ea'), latest_ga ('ga'), latest_sts ('sts'), latest_mts ('mts'), latest_lts ('lts')",
		},
		&cli.BoolFlag{
			Name:    optionMaintained,
			Aliases: []string{"mt"},
			Usage:   "Maintained or not.",
		},
		&cli.BoolFlag{
			Name:    optionEarlyAccess,
			Aliases: []string{"ea"},
			Usage:   "Early Access.",
		},
		&cli.BoolFlag{
			Name:    optionGeneralAvailability,
			Aliases: []string{"ga"},
			Usage:   "General availability.",
		},
	},
}

//majorVersionsStruct defines the json structure which is replied for /major_versions from REST.
type majorVersionsStruct []struct {
	MajorVersion  int      `json:"major_version"`
	TermOfSupport string   `json:"term_of_support"`
	Maintained    bool     `json:"maintained"`
	Versions      []string `json:"versions"`
}

//majorVersionsLatestStruct defines the json structure which will be replied by /major_versions/latest_..
type majorVersionsLatestStruct struct {
	MajorVersion  int      `json:"major_version"`
	TermOfSupport string   `json:"term_of_support"`
	Maintained    bool     `json:"maintained"`
	Versions      []string `json:"versions"`
}

func actionMajorVersions(ctx *cli.Context) error {
	var url = fmt.Sprintf("%s/major_versions", foojayBaseAPI)

	if ctx.IsSet(optionLatest) {
		latest, err := fromShortToLatest(ctx.String(optionLatest))
		check.IfError(err)
		majorVersionQuery(fmt.Sprintf("%s/%s", url, latest))
	} else if ctx.IsSet(optionMajorVersion) {

		if !(ctx.IsSet(optionEarlyAccess) != ctx.IsSet(optionGeneralAvailability)) {
			return fmt.Errorf("either --ea or --ga must be given")
		}

		givenVersion := ctx.String(optionMajorVersion)
		url = fmt.Sprintf("%s/%s", url, givenVersion)
		if ctx.IsSet(optionEarlyAccess) {
			url = fmt.Sprintf("%s/ea", url)
		}
		if ctx.IsSet(optionGeneralAvailability) {
			url = fmt.Sprintf("%s/ga", url)
		}
		fmt.Printf("URL: %s\n", url)

		majorVersionQuery(url)
	} else if ctx.IsSet(optionMaintained) || ctx.IsSet(optionEarlyAccess) || ctx.IsSet(optionGeneralAvailability) {
		query := []string{}
		if ctx.IsSet(optionMaintained) {
			query = append(query, "maintained=true")
		}
		if ctx.IsSet(optionEarlyAccess) {
			query = append(query, "ea=true")
		}
		if ctx.IsSet(optionGeneralAvailability) {
			query = append(query, "ga=true")
		}
		url = fmt.Sprintf("%s?%s", url, strings.Join(query, "&"))
		fmt.Printf("URL: %s\n", url)
		majorVersionMaintainedEaGa(url)
	} else {
		fmt.Printf("URL: %s\n", url)
		majorVersion(url)

	}

	return nil
}

func majorVersion(url string) {
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
	table.Render()
}

func majorVersionMaintainedEaGa(url string) {
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
	table.Render()
}

func majorVersionQuery(url string) {
	var majorLatest majorVersionsLatestStruct
	helper.GetData(url, &majorLatest)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Major Version", "Maintained", "Term of Support", "Versions"})
	table.SetAutoWrapText(true)
	table.SetRowLine(true)

	majorVersion := fmt.Sprintf("%d", majorLatest.MajorVersion)
	maintained := fromBoolToYesNo(majorLatest.Maintained)
	versions := strings.Join(majorLatest.Versions, ", ")
	row := []string{majorVersion, maintained, majorLatest.TermOfSupport, versions}
	table.Append(row)
	table.Render()
}

//fromBoolToYesNo Will convert true => 'Yes' and 'No' otherwise.
func fromBoolToYesNo(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

//fromShortToLatest Will convert 'ea', 'ga', 'sts', 'mts', 'lts' into `latest_...`.
func fromShortToLatest(value string) (result string, err error) {
	switch strings.ToLower(value) {
	case "ea":
		fallthrough
	case "ga":
		fallthrough
	case "sts":
		fallthrough
	case "mts":
		fallthrough
	case "lts":
		result = fmt.Sprintf("latest_%s", value)
		err = nil
		break
	default:
		result = ""
		err = fmt.Errorf("an invalid value '%s' given only one of following is valid: 'ea','ga','sts','mts' or 'lts'", value)
	}
	return result, err
}
