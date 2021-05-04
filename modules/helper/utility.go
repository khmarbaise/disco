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

	defer response.Body.Close()
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

	req.Header.Set("user-agent", fmt.Sprintf("disco go command line utility version: %s", Version))
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	return response, err
}
