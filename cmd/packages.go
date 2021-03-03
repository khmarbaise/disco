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

const (
	optionVersionByDefinition  = "version_by_definition"
	optionArchitecture         = "architecture"
	optionDistribution         = "distribution"
	optionArchiveType          = "archive_type"
	optionPackageType          = "package_type"
	optionOperatingSystem      = "operating_system"
	optionLibcType             = "libc_type"
	optionReleaseStatus        = "release_status"
	optionTermOfSupport        = "term_of_support"
	optionBitness              = "bitness"
	optionJavaFXBundled        = "javafx_bundled"
	optionDirectlyDownloadable = "directly_downloadable"
	optionFromVersion          = "from_version"
	optionToVersion            = "to_version"
)

//Packages ....
var Packages = cli.Command{
	Name:        "packages",
	Aliases:     []string{"pkg"},
	Usage:       "packages on issue",
	Description: "packages ...",
	Action:      packages,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    optionVersion,
			Aliases: []string{"v"},
			Usage:   "for example 1.8.0_265 or 11 or 13.0.5.1",
		},
		&cli.StringFlag{
			Name:    optionFromVersion,
			Aliases: []string{"fv"},
			Usage:   "for example 1.8.0_265 or 11 or 13.0.5.1",
		},
		&cli.StringFlag{
			Name:    optionToVersion,
			Aliases: []string{"tv"},
			Usage:   "for example 1.8.0_265 or 11 or 13.0.5.1",
		},
		&cli.StringFlag{
			Name:    optionVersionByDefinition,
			Aliases: []string{"vbd"},
			Usage:   "The version will be calculated from the given parameter (latest, latest_sts, latest_mts, latest_lts)",
		},
		&cli.StringFlag{
			Name:    optionArchitecture,
			Aliases: []string{"arch"},
			Usage:   "Architecture for example aarch64, arm, arm64, mips, ppc, ppc64, ppc64le, riscv64, s390x, sparc, sparcv9, x64, x86, amd64.",
		},
		&cli.StringFlag{
			Name:    optionDistribution,
			Aliases: []string{"distro", "dist"},
			Usage:   "Distribution for example  'aoj', 'aoj_openj9', 'dragonwell', 'corretto', 'liberica', 'oracle_open_jdk', 'redhat', 'sap_machine', 'zulu'.",
		},
		&cli.StringFlag{
			Name:    optionArchiveType,
			Aliases: []string{"ext"},
			Usage:   "File extension e.g. 'cab', 'deb', 'dmg', 'exe', 'msi', 'pkg', 'rpm', 'tar', 'zip'",
		},
		&cli.StringFlag{
			Name:    optionPackageType,
			Aliases: []string{"pt"},
			Usage:   "Package type like 'jre', 'jdk'",
		},
		&cli.StringFlag{
			Name:    optionOperatingSystem,
			Aliases: []string{"os"},
			Usage:   "Operating System for example 'windows', 'macos', 'linux'.",
		},
		&cli.StringFlag{
			Name:    optionLibcType,
			Aliases: []string{"lt"},
			Usage:   "Type of libc for example 'glibc', 'libc', 'musl', 'c_std_lib'.",
		},
		&cli.StringFlag{
			Name:    optionReleaseStatus,
			Aliases: []string{"rs"},
			Usage:   "The release status early access or general availability ('ea', 'ga').",
		},
		&cli.StringFlag{
			Name:    optionTermOfSupport,
			Aliases: []string{"tos"},
			Usage:   "Term of support for example 'sts' (short term support), 'mts' (mid term support), 'lts' (long term stable).",
		},
		&cli.StringFlag{
			Name:    optionBitness,
			Aliases: []string{"b"},
			Usage:   "Bitness for example '32' or '64' bits.",
		},

		&cli.BoolFlag{
			Name:    optionJavaFXBundled,
			Aliases: []string{"fx"},
			Usage:   "With JavaFX",
		},
		&cli.BoolFlag{
			Name:    optionDirectlyDownloadable,
			Aliases: []string{"dd"},
			Usage:   "Directly downloadable.",
		},
		//TODO: I think two separate options are better.
		&cli.StringFlag{
			Name:  "latest",
			Usage: "Latest for example 'overall', 'per_distro'",
		},
		&cli.BoolFlag{
			Name:  optionVerbose,
			Usage: "Printout all versions.",
		},
	},
}

//PackagesStructure Defines structure for REST API /packages.
type PackagesStructure []struct {
	ID                   string `json:"id"`
	ArchiveType          string `json:"archive_type"`
	Distribution         string `json:"distribution"`
	MajorVersion         int    `json:"major_version"`
	JavaVersion          string `json:"java_version"`
	DistributionVersion  string `json:"distribution_version"`
	LatestBuildAvailable bool   `json:"latest_build_available"`
	ReleaseStatus        string `json:"release_status"`
	TermOfSupport        string `json:"term_of_support"`
	OperatingSystem      string `json:"operating_system"`
	LibCType             string `json:"lib_c_type"`
	Architecture         string `json:"architecture"`
	PackageType          string `json:"package_type"`
	JavafxBundled        bool   `json:"javafx_bundled"`
	DirectlyDownloadable bool   `json:"directly_downloadable"`
	Filename             string `json:"filename"`
	EphemeralID          string `json:"ephemeral_id"`
}

func packages(ctx *cli.Context) error {
	var url = fmt.Sprintf("%s/packages", foojayBaseAPI)
	query := []string{}

	if ctx.IsSet(optionVersion) {
		query = append(query, fmt.Sprintf("version=%s", ctx.String(optionVersion)))
	}
	if ctx.IsSet(optionFromVersion) {
		query = append(query, fmt.Sprintf("from_version=%s", ctx.String(optionFromVersion)))
	}
	if ctx.IsSet(optionToVersion) {
		query = append(query, fmt.Sprintf("to_version=%s", ctx.String(optionToVersion)))
	}
	if ctx.IsSet(optionVersionByDefinition) {
		query = append(query, fmt.Sprintf("version_by_definition=%s", ctx.String(optionVersionByDefinition)))
	}
	if ctx.IsSet(optionLibcType) {
		query = append(query, fmt.Sprintf("libc_type=%s", ctx.String(optionLibcType)))
	}
	if ctx.IsSet(optionArchiveType) {
		query = append(query, fmt.Sprintf("archive_type=%s", ctx.String(optionArchiveType)))
	}
	if ctx.IsSet(optionArchitecture) {
		query = append(query, fmt.Sprintf("architecture=%s", ctx.String(optionArchitecture)))
	}
	if ctx.IsSet(optionDistribution) {
		query = append(query, fmt.Sprintf("distro=%s", ctx.String(optionDistribution)))
	}
	if ctx.IsSet(optionBitness) {
		query = append(query, fmt.Sprintf("bitness=%s", ctx.String(optionBitness)))
	}
	if ctx.IsSet(optionPackageType) {
		query = append(query, fmt.Sprintf("package_type=%s", ctx.String(optionPackageType)))
	}
	if ctx.IsSet(optionOperatingSystem) {
		query = append(query, fmt.Sprintf("operating_system=%s", ctx.String(optionOperatingSystem)))
	}
	if ctx.IsSet(optionReleaseStatus) {
		query = append(query, fmt.Sprintf("release_status=%s", ctx.String(optionReleaseStatus)))
	}
	if ctx.IsSet(optionTermOfSupport) {
		query = append(query, fmt.Sprintf("support_term=%s", ctx.String(optionTermOfSupport)))
	}
	if ctx.IsSet(optionDirectlyDownloadable) {
		query = append(query, fmt.Sprintf("directly_downloadable=%s", helper.FromBoolToYesNo(ctx.Bool(optionDirectlyDownloadable))))
	}
	if ctx.IsSet(optionJavaFXBundled) {
		query = append(query, fmt.Sprintf("javafx_bundled=%s", helper.FromBoolToYesNo(ctx.IsSet(optionJavaFXBundled))))
	}

	if len(query) > 0 {
		url = fmt.Sprintf("%s?%s", url, strings.Join(query, "&"))
	}
	fmt.Printf("URL: %s\n", url)

	var packagesStructure = PackagesStructure{}
	helper.GetData(url, &packagesStructure)

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		//"ID",
		"ArchiveType",
		"Distro",
		"Version",
		"JavaVersion",
		"DistributionVersion",
		"Build",
		"Status",
		"Support",
		"OS",
		"LibCType",
		"Arch",
		"PT",
		"FX",
		"Downloadable",
		//"Filename",
		//"EphemeralID",
	})
	table.SetAutoWrapText(true)
	table.SetRowLine(true)

	for _, v := range packagesStructure {
		row := []string{
			//fmt.Sprintf("%s", v.ID),
			v.ArchiveType,
			v.Distribution,
			fmt.Sprintf("%d", v.MajorVersion),
			v.JavaVersion,
			v.DistributionVersion,
			helper.FromBoolToYesNo(v.LatestBuildAvailable),
			v.ReleaseStatus,
			v.TermOfSupport,
			v.OperatingSystem,
			v.LibCType,
			v.Architecture,
			v.PackageType,
			helper.FromBoolToYesNo(v.JavafxBundled),
			helper.FromBoolToYesNo(v.DirectlyDownloadable),
			//v.Filename,
			//v.EphemeralID,
		}
		table.Append(row)
	}
	table.Render()
	return nil
}
