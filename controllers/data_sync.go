package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DataSyncResponse struct {
	Status     string `json:"status"`
	HasNewData bool   `json:"hasNewData"`
	Code       string `json:"code"`
}

func SyncAccountData(code string) DataSyncResponse {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/sync", id)
	req, _ := http.NewRequest("POST", url, nil)
	mono_client := http.Client{}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to sync account data"))
	}

	b, _ := ioutil.ReadAll(resp.Body)

	syncResponse := DataSyncResponse{}
	err = json.Unmarshal(b, &syncResponse)

	if err != nil {
		panic(err)
	}
	return syncResponse
}

type ReAuthorizationResponse struct {
	Token string `json:"token"`
}

func ReAuthorizeAccount(code string) ReAuthorizationResponse {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/reauthorise", id)
	req, _ := http.NewRequest("POST", url, nil)
	mono_client := http.Client{}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to sync account data"))
	}

	b, _ := ioutil.ReadAll(resp.Body)
	reAuthResponse := ReAuthorizationResponse{}
	json.Unmarshal(b, &reAuthResponse)

	return reAuthResponse
}
