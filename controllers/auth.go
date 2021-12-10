package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetAuthenticationDetails(code string) map[string]string {
	mono_client := http.Client{}
	req_body := make(map[string]interface{})
	req_body["code"] = code
	json_req, err := json.Marshal(req_body)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(json_req)

	url := "https://api.withmono.com/account/auth"
	req, err := http.NewRequest("POST", url, b)
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get authentication details"))
	}
	resp_body, _ := ioutil.ReadAll(resp.Body)

	output := make(map[string]string)

	_ = json.Unmarshal(resp_body, &output)
	return output
}
