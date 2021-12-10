package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Institution struct {
	Name     string `json:"name"`
	BankCode string `json:"bankCode"`
	Type     string `json:"type"`
}

type Account struct {
	ID            string      `json:"_id"`
	Institution   Institution `json:"institution"`
	Name          string      `json:"name"`
	AccountNumber string      `json:"accountNumber"`
	Type          string      `json:"type"`
	Balance       int         `json:"balance"`
	Currency      string      `json:"currency"`
	Bvn           string      `json:"bvn"`
}

type Meta struct {
	DataStatus string `json:"data_status"`
	AuthMethod string `json:"auth_method"`
}

type FullAccountDetails struct {
	Meta    Meta    `json:"meta"`
	Account Account `json:"account"`
}

type Statement struct {
	ID        string    `json:"_id"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date"`
	Narration string    `json:"narration"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
}

type AccountStatements struct {
	StatementsMeta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Statements []Statement `json:"data"`
}

func GetFullAccoutDetails(code string) FullAccountDetails {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%v", id)
	mono_client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	output := FullAccountDetails{}
	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get full account details"))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &output)
	if err != nil {
		panic(err)
	}
	return output
}

func GetAccountStatements(code string) AccountStatements {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%v/statement", id)
	mono_client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")

	resp, err := mono_client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to get all transactions"))
	}

	transactionHistory := AccountStatements{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &transactionHistory)
	if err != nil {
		panic(err)
	}

	return transactionHistory
}

type Paging struct {
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type Transaction struct {
	ID        string    `json:"_id"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date"`
	Narration string    `json:"narration"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
}
type TransactionHistory struct {
	Paging       Paging        `json:"paging"`
	Transactions []Transaction `json:"data"`
}

func GetAllTransactions(code string) TransactionHistory {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/transactions", id)
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

	transactionsHistory := TransactionHistory{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &transactionsHistory)

	if err != nil {
		panic(err)
	}
	return transactionsHistory
}

type AccountIncome struct {
	Type       string  `json:"type"`
	Amount     int     `json:"amount"`
	Employer   string  `json:"employer"`
	Confidence float64 `json:"confidence"`
}

func GetAccountIncome(code string) AccountIncome {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/transactions", id)
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

	accountIncome := AccountIncome{}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &accountIncome)

	if err != nil {
		panic(err)
	}
	return accountIncome

}

type AccountIdentity struct {
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Gender        string `json:"gender"`
	DataOfBirth   string `json:"dob"`
	BVN           string `json:"bvn"`
	MaritalStatus string `json:"maritalStatus"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
}

func GetAccountIdentity(code string) AccountIdentity {
	id := GetAuthenticationDetails(code)["id"]
	url := fmt.Sprintf("https://api.withmono.com/accounts/%s/indentity", id)
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

	identity := AccountIdentity{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &identity)

	if err != nil {
		panic(err)
	}
	return identity
}
