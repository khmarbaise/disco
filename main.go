// Copyright 2021 The Disco Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// disco is command line tool to support the handling of Foojay DISCO REST API
package main

import (
	"fmt"
	"github.com/khmarbaise/disco/cmd"
	"github.com/khmarbaise/disco/modules/helper"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "disco"
	app.Usage = "Command line tool to explore the Foojay Discovery API"
	app.Description = "A command line based Foojay Discovery API (disco) tool to explore the information of the API"
	app.Version = helper.Version + formatBuiltWith(helper.Tags)
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&cmd.Distributions,
		&cmd.MajorVersions,
		&cmd.Packages,
		&cmd.EphemeralIds,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}

}

func formatBuiltWith(Tags string) string {
	if len(Tags) == 0 {
		return ""
	}

	return " built with: " + strings.Replace(Tags, " ", ", ", -1)
}
