package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetAllDetailsWithBVN(bvn string) (map[string]interface{}, error) {
	mono_client := http.Client{}
	url := "https://api.withmono.com/360view"
	data := make(map[string]interface{})
	data["bvn"] = bvn
	b, _ := json.Marshal(data)
	body := bytes.NewBuffer(b)

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")
	req.Header.Add("Content-Type", "application/json")

	resp, err := mono_client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get users bvn"))
	}

	b, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	output := make(map[string]interface{})
	err = json.Unmarshal(b, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
