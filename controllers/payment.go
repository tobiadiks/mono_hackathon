package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func generateReference() string {
	keys := []string{"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z", "0", "1",
		"2", "3", "4", "5", "6", "7", "8", "9"}

	reference := []byte{}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 11; i++ {
		reference = append(reference, reference[rand.Intn(len(keys))])

	}
	return string(reference)
}

type PaymentInitializationRequest struct {
	Amount      string `json:"amount"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Reference   string `json:"reference"`
	Account     string `json:"string"`
}

type PaymentInitializationResponse struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Reference   string    `json:"reference"`
	PaymentLink string    `json:"payment_link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func InitializeOneTimePayment(code string, amount int64, description string) PaymentInitializationResponse {
	// generate payment reference
	reference := generateReference()
	id := GetAuthenticationDetails(code)["id"]

	paymentRequest := PaymentInitializationRequest{
		Amount:      strconv.Itoa(int(amount * 100)),
		Type:        "onetime-debit",
		Description: description,
		Reference:   reference,
		Account:     id,
	}

	url := "https://api.withmono.com/v1/payments/initiate"
	json_req, err := json.Marshal(paymentRequest)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer(json_req)
	req, _ := http.NewRequest("POST", url, b)

	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")
	req.Header.Add("Content-Type", "application/json")

	mono_client := http.Client{}

	resp, _ := mono_client.Do(req)

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to initiale payment"))
	}

	paymentResponse := PaymentInitializationResponse{}
	response_body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response_body, &paymentResponse)

	if err != nil {
		panic(err)
	}

	return paymentResponse
}

type VerifyPaymentRequest struct {
	Reference string `json:"reference"`
}

type PaymentVerificationResponse struct {
	Type string `json:"type"`
	Data struct {
		MainID      string    `json:"_id"`
		ID          string    `json:"id"`
		Status      string    `json:"status"`
		Amount      int       `json:"amount"`
		Description string    `json:"description"`
		Fee         int       `json:"fee"`
		Currency    string    `json:"currency"`
		Account     string    `json:"account"`
		Customer    string    `json:"customer"`
		Reference   string    `json:"reference"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"data"`
}

func VerifyPayment(reference string) PaymentVerificationResponse {
	url := "https://api.withmono.com/v1/payments/verify"
	verifyPaymentRequest := VerifyPaymentRequest{
		Reference: reference,
	}
	json_req, _ := json.Marshal(verifyPaymentRequest)
	b := bytes.NewBuffer(json_req)

	req, _ := http.NewRequest("POST", url, b)

	req.Header.Add("mono-sec-key", "test_sk_qMyxK2nzsIk8hIHX14dC")
	req.Header.Add("Content-Type", "application/json")

	mono_client := http.Client{}

	resp, _ := mono_client.Do(req)

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("failed to initiale payment"))
	}

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	paymentVerificationResponse := PaymentVerificationResponse{}
	json.Unmarshal(resp_body, &paymentVerificationResponse)

	return paymentVerificationResponse

}
