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
			Usage:   "Release status for example 'ea', 'ga'.",
		},
		&cli.StringFlag{
			Name:    "term_of_support",
			Aliases: []string{"tos"},
			Usage:   "Term of support for example 'sts', 'mts', 'lts'.",
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

func packages(ctx *cli.Context) error {

	fmt.Println("Not yet implemented.")
	return nil
}
