package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AssetDetails struct {
	Symbol string `json:"symbol"`
	Price  int    `json:"price"`
}

type Asset struct {
	ID       string       `json:"_id"`
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Cost     int          `json:"cost"`
	Return   int          `json:"return"`
	Quantity float64      `json:"quantity"`
	Currency string       `json:"currency"`
	Details  AssetDetails `json:"details"`
}

type Assets []Asset

func GetAccountAssets(code string) Assets {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/assets", id)
	req, _ := http.NewRequest("GET", url, nil)
	mono_client := http.Client{}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get all transactions"))
	}

	assets := Assets{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &assets)

	if err != nil {
		panic(err)
	}
	return assets
}

type EarningAsset struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	SalePrice    int    `json:"sale_price"`
	QuantitySold int    `json:"quantity_sold"`
}

type Earning struct {
	ID        string       `json:"_id"`
	Amount    int          `json:"amount"`
	Narration string       `json:"narration"`
	Date      time.Time    `json:"date"`
	Asset     EarningAsset `json:"asset"`
}

type Earnings []Earning

func GetAllAccountEarnings(code string) Earnings {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/earings", id)
	req, _ := http.NewRequest("GET", url, nil)
	mono_client := http.Client{}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get all transactions"))
	}

	earnings := Earnings{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &earnings)

	if err != nil {
		panic(err)
	}
	return earnings
}
