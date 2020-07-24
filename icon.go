package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Icon struct {
	Title  string `json:"title"`
	Hex    string `json:"hex"`
	Source string `json:"source"`
}

func getIcons() ([]Icon, error) {
	res, err := http.DefaultClient.Get("https://raw.githubusercontent.com/simple-icons/simple-icons/develop/_data/simple-icons.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	defer io.Copy(ioutil.Discard, res.Body)

	var simpleIcons struct {
		Icons []Icon `json:"icons"`
	}

	err = json.NewDecoder(res.Body).Decode(&simpleIcons)
	if err != nil {
		return nil, err
	}

	return simpleIcons.Icons, nil
}
