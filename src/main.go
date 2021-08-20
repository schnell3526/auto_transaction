package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type apiKey struct {
	API_KEY    string
	SECRET_KEY string
}

func main() {
	raw, err := ioutil.ReadFile("../.config")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var key apiKey

	json.Unmarshal(raw, &key)

	fmt.Printf("API_KEY = %s\n", key.API_KEY)
	fmt.Printf("SECRET_KEY = %s\n", key.SECRET_KEY)
}
