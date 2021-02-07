// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
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
			Name:     "version",
			Aliases:  []string{"v"},
			Usage:    "for example 1.8.0_265 or 11 or 13.0.5.1",
			Required: true,
		},
		// "from_version" The packages where the version is larger than from_version (e.g. 11.0.1)
		// "to_version" The packages where the version is smaller than to_version (e.g. 11.0.5)
		&cli.StringFlag{
			Name:    "version_by_definition",
			Aliases: []string{"vbd"},
			Usage:   "The version will be calculated from the given parameter (latest, latest_sts, latest_mts, latest_lts)",
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Define the distribution name for example 'zulu', 'oracle'.",
		},
		&cli.StringFlag{
			Name:    "architecture",
			Aliases: []string{"arch"},
			Usage:   "Architecture for example aarch64, arm, arm64, mips, ppc, ppc64, ppc64le, riscv64, s390x, sparc, sparcv9, x64, x86, amd64.",
		},
		&cli.StringFlag{
			Name:    "distribution",
			Aliases: []string{"distro", "dist"},
			Usage:   "Distribution for example  'aoj', 'aoj_openj9', 'dragonwell', 'corretto', 'liberica', 'oracle_open_jdk', 'redhat', 'sap_machine', 'zulu'.",
		},
		&cli.StringFlag{
			Name:    "archive_type",
			Aliases: []string{"ext"},
			Usage:   "File extension e.g. 'cab', 'deb', 'dmg', 'exe', 'msi', 'pkg', 'rpm', 'tar', 'zip'",
		},
		&cli.StringFlag{
			Name:    "package_type",
			Aliases: []string{"pt"},
			Usage:   "Package type like 'jre', 'jdk'",
		},
		&cli.StringFlag{
			Name:    "operating_system",
			Aliases: []string{"os"},
			Usage:   "Operating System for example 'windows', 'macos', 'linux'.",
		},
		&cli.StringFlag{
			Name:    "libc_type",
			Aliases: []string{"lt"},
			Usage:   "Type of libc for example 'glibc', 'libc', 'musl', 'c_std_lib'.",
		},
		&cli.StringFlag{
			Name:    "release_status",
			Aliases: []string{"rs"},
			Usage:   "The release status early access or general availability ('ea', 'ga').",
		},
		&cli.StringFlag{
			Name:    "term_of_support",
			Aliases: []string{"tos"},
			Usage:   "Term of support for example 'sts' (short term support), 'mts' (mid term support), 'lts' (long term stable).",
		},
		&cli.StringFlag{
			Name:    "bitness",
			Aliases: []string{"b"},
			Usage:   "Bitness for example '32' or '64' bits.",
		},

		&cli.BoolFlag{
			Name:    "javafx_bundled",
			Aliases: []string{"fx"},
			Usage:   "With JavaFX",
		},
		&cli.BoolFlag{
			Name:    "directly_downloadable",
			Aliases: []string{"dd"},
			Usage:   "Directly downloadable.",
		},
		//TODO: I think two separate options are better.
		&cli.StringFlag{
			Name:  "latest",
			Usage: "Latest for example 'overall', 'per_distro'",
		},
		&cli.BoolFlag{
			Name:  "verbose",
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

	fmt.Println("Not yet implemented.")
	return nil
}
