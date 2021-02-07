// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//EphemeralIds Describe
var EphemeralIds = cli.Command{
	Name:        "ephemeralids",
	Aliases:     []string{"eids", "ids"},
	Usage:       "ephemeralids .....",
	Description: "ephemeralids ....descritpion",
	Action:      ephemeralIds,
}

//AutoGenerated Result while calling REST API
type AutoGenerated struct {
	Filename          string `json:"filename"`
	DirectDownloadURI string `json:"direct_download_uri"`
	DownloadSiteURI   string `json:"download_site_uri"`
}

func ephemeralIds(ctx *cli.Context) error {

	var url = fmt.Sprintf("%s/ephemeral_ids", FoojayBaseAPI)
	fmt.Printf("URL: %s\n", url)

	return nil
}
