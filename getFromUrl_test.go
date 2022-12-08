package qboauth

import (
	"io/ioutil"
	"os"
)

func (c mockClient) getFromUrl() (Document, error) {
	var jsonFile *os.File
	var err error
	if c.env == Production {
		jsonFile, err = os.Open("./testdata/production.json")
	} else {
		jsonFile, err = os.Open("./testdata/sandbox.json")
	}
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	return parseJsonIntoStruct(byteValue, c.env)
}
