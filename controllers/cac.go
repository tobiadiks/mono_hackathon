package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Business struct {
	State            interface{} `json:"state"`
	ID               int         `json:"id"`
	Address          string      `json:"address"`
	ApprovedName     string      `json:"approvedName"`
	RcNumber         string      `json:"rcNumber"`
	BranchAddress    string      `json:"branchAddress"`
	RegistrationDate time.Time   `json:"registrationDate"`
	ClassificationID int         `json:"classificationId"`
	Email            interface{} `json:"email"`
	Lga              interface{} `json:"lga"`
	City             interface{} `json:"city"`
	Status           string      `json:"status"`
}

type BusinessDetailsRequest struct {
	Name string `json:"name"`
}

func GetBusinessDetails(name string) Business {
	url := "https://api.withmono.com/v1/cac/lookup"
	businessDetailsRequest := BusinessDetailsRequest{Name: name}

	req_body, _ := json.Marshal(businessDetailsRequest)
	b := bytes.NewBuffer(req_body)
	req, _ := http.NewRequest("POST", url, b)

	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")
	req.Header.Add("Content-Type", "application/json")

	mono_client := http.Client{}

	resp, _ := mono_client.Do(req)

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get business details"))
	}

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	business := Business{}
	json.Unmarshal(resp_body, &business)
	return business
}
