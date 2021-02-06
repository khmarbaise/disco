package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
