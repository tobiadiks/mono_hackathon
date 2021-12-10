package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type BusinessInstitution struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Coverage struct {
		Countries []string `json:"countries"`
		Business  bool     `json:"business"`
		Personal  bool     `json:"personal"`
	} `json:"coverage"`
	Products []string    `json:"products"`
	Website  interface{} `json:"website"`
}

type BusinessInstitutions []BusinessInstitution

func GetAllBusinessInstutions() BusinessInstitutions {
	url := "https://api.withmono.com/coverage"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	mono_client := http.Client{}

	resp, _ := mono_client.Do(req)

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get businesses "))
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	businesses := BusinessInstitutions{}
	json.Unmarshal(b, &businesses)
	return businesses
}
