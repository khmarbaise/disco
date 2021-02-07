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
	response, err := http.Get(checkURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//Need to think about this.
	if response.StatusCode != http.StatusOK {
		fmt.Printf("%s\n", response.Status)
		os.Exit(2)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &v)
}
