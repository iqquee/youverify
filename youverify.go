package youverify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	// SandboxBaseURL is the test environment base url
	SandboxBaseURL string = "https://api.sandbox.youverify.co/"
	// ProductionBaseURL is the live environment base url
	ProductionBaseURL string = "https://api.youverify.co/"

	// StatusCode200 when the server returns a status code 200
	StatusCode200 StatusCode = "[Successful] - Your request was successful."
	// StatusCode400 when the server returns a status code 400
	StatusCode400 StatusCode = "[Bad request] - Most likely an invalid syntax. Check all parameters."
	// StatusCode401 when the server returns a status code 401
	StatusCode401 StatusCode = "[Unauthorized] - You are unauthorized for this request. Contact support@youverify.co."
	// StatusCode402 when the server returns a status code 402
	StatusCode402 StatusCode = "[Payment Required] - You balance is most likely low and you are required to top up your balance."
	// StatusCode404 when the server returns a status code 404
	StatusCode404 StatusCode = "[Not found] - URL not recognized. Check to confirm the right URL."
	// StatusCode405 when the server returns a status code 405
	StatusCode405 StatusCode = "[Method not found] - The request is disabled. Check the URL or rollback recent upgrades."
	// StatusCode408 when the server returns a status code 408
	StatusCode408 StatusCode = "[Request timeout] - Your request took longer than it should have. Check your internet connection."
	// StatusCode424 when the server returns a status code 424
	StatusCode424 StatusCode = "[Failed Dependency] - Third Party service Failure."
	// StatusCode429 when the server returns a status code 429
	StatusCode429 StatusCode = "[Too many requests] - You have sent too many requests that has exceeded the rate limit. You need to wait a while."
	// StatusCode500 when the server returns a status code 500
	StatusCode500 StatusCode = "[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co"
	// StatusCode502 when the server returns a status code 502
	StatusCode502 StatusCode = "[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co"
	// StatusCode503 when the server returns a status code 503
	StatusCode503 StatusCode = "[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co"
	// StatusCode504 when the server returns a status code 504
	StatusCode504 StatusCode = "[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co"
)

type (
	// StatusCode type
	StatusCode string

	// Kyc for the KYC (Know Your Customer)
	Kyc struct{}
	// Kyb for the KYB (Know Your Business)
	Kyb struct{}
	// RiskInteligence for the risk inteligence
	RiskInteligence struct{}
	// AML for the AML (Anti-Money Laundry)
	AML struct{}

	youverify struct {
		Http        *http.Client
		BaseURL     string
		SecretToken string
	}

	InitializePayload struct {
		IsLive      bool
		SecretToken string
		Http        *http.Client
	}

	Client struct {
		Kyc             Kyc
		Kyb             Kyb
		youverify       youverify
		RiskInteligence RiskInteligence
		AMl             AML
	}
)

func Initialize(i InitializePayload) *Client {
	var baseUrl string
	if i.IsLive {
		baseUrl = ProductionBaseURL
	} else {
		baseUrl = SandboxBaseURL
	}

	return &Client{
		youverify: youverify{
			BaseURL:     baseUrl,
			SecretToken: i.SecretToken,
			Http:        i.Http,
		},
	}
}

/*
newRequest makes a http request to the urlbox server and decodes the server response into the reqBody parameter passed into the newRequest method
*/
func (c *Client) newRequest(method, reqURL string, reqBody interface{}) ([]byte, int, error) {
	newURL := c.youverify.BaseURL + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return nil, 0, errors.Wrap(err, "http client ::: unable to marshal request struct")
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	bearer := fmt.Sprintf("Bearer %v", c.youverify.BaseURL)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	if err != nil {
		return nil, 0, errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := c.youverify.Http.Do(req)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http client ::: client failed to execute request")
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http client ::: client failed to read file")
	}

	return b, res.StatusCode, nil
}
