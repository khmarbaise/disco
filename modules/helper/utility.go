// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//GetData Will get on the given URL and convert the JSON into the given datastructure.
func GetData(checkURL string, v interface{}) {
	response, err := privateGet(checkURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	//Need to think about this.
	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Status Code: %s\n", response.Status)
		os.Exit(2)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close();
	json.Unmarshal(responseData, &v)
}

//FromBoolToYesNo Convert from bool to string.
func FromBoolToYesNo(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

//privateGet This method will add a http-header which contains the user-agent.
func privateGet(checkURL string) (result *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", checkURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	req.Header.Set("user-agent", "disco go command line utility version: " + Version)
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	return response, err
}
