package youverify

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	// SandboxBaseURL is the test environment base url
	SandboxBaseURL string = "https://api.sandbox.youverify.co/v2/api/"
	// ProductionBaseURL is the live environment base url
	ProductionBaseURL string = "https://api.youverify.co/v2/api/"
	// methodPOST for an http POST method
	methodPOST string = "POST"
	// methodGET for an http GET method
	methodGET string = "GET"
)

var (
	// baseURL holds the value of the baseURL
	baseURL string
	// secretToken holds the value of the secretToken
	secretToken string
	// client holds the http client
	client *http.Client

	// ErrStatusCode200 when the server returns a status code 200
	ErrStatusCode200 StatusCode = errors.New("[Successful] - Your request was successful.")
	// ErrStatusCode400 when the server returns a status code 400
	ErrStatusCode400 StatusCode = errors.New("[Bad request] - Most likely an invalid syntax. Check all parameters.")
	// ErrStatusCode401 when the server returns a status code 401
	ErrStatusCode401 StatusCode = errors.New("[Unauthorized] - You are unauthorized for this request. Contact support@youverify.co.")
	// ErrStatusCode402 when the server returns a status code 402
	ErrStatusCode402 StatusCode = errors.New("[Payment Required] - You balance is most likely low and you are required to top up your balance.")
	// ErrStatusCode404 when the server returns a status code 404
	ErrStatusCode404 StatusCode = errors.New("[Not found] - URL not recognized. Check to confirm the right URL.")
	// ErrStatusCode405 when the server returns a status code 405
	ErrStatusCode405 StatusCode = errors.New("[Method not found] - The request is disabled. Check the URL or rollback recent upgrades.")
	// ErrStatusCode408 when the server returns a status code 408
	ErrStatusCode408 StatusCode = errors.New("[Request timeout] - Your request took longer than it should have. Check your internet connection.")
	// ErrStatusCode424 when the server returns a status code 424
	ErrStatusCode424 StatusCode = errors.New("[Failed Dependency] - Third Party service Failure.")
	// ErrStatusCode429 when the server returns a status code 429
	ErrStatusCode429 StatusCode = errors.New("[Too many requests] - You have sent too many requests that has exceeded the rate limit. You need to wait a while.")
	// ErrStatusCode500 when the server returns a status code 500
	ErrStatusCode500 StatusCode = errors.New("[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co")
	// ErrStatusCode502 when the server returns a status code 502
	ErrStatusCode502 StatusCode = errors.New("[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co")
	// ErrStatusCode503 when the server returns a status code 503
	ErrStatusCode503 StatusCode = errors.New("[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co")
	// ErrStatusCode504 when the server returns a status code 504
	ErrStatusCode504  StatusCode = errors.New("[Server error] - This is a very rare occurrence where the server is unable to process a request properly. Contact support@youverify.co")
	errSubjectConsent            = errors.New("Subject consent field must be true")
)

type (
	// StatusCode type
	StatusCode error

	// Client default client for youverify
	Client struct{}
	// Kyc for the KYC (Know Your Customer)
	Kyc struct{}
	// Kyb for the KYB (Know Your Business)
	Kyb struct{}
	// RiskInteligence for the risk inteligence
	RiskInteligence struct{}
	// AML for the AML (Anti-Money Laundry)
	AML struct{}

	// InitializeRequest is the payload that initializes youverify
	InitializeRequest struct {
		IsLive      bool
		SecretToken string
		Http        *http.Client
	}

	Nigeria struct{}
)

/*
Initialize is a function that sets up youverify configs.

This function takes in three(3) paramaters: Http - which is the http.Client{} for making requests, IsLive - is a bool which specifies the environment and SecretToken - is the secretToken used to authorize requests.
*/
func Initialize(Http *http.Client, IsLive bool, SecretToken string) *Client {
	// var baseUrl string
	if IsLive {
		baseURL = ProductionBaseURL
	} else {
		baseURL = SandboxBaseURL
	}

	secretToken = SecretToken
	client = Http
	return &Client{}
}

// KYC returns methods for KYC (Know Your Customer)
func (c *Client) KYC() *Kyc {
	return &Kyc{}
}

// Nigeria returns methods for KYC (Know Your Customer) in Nigeria
func (c *Kyc) Nigeria() *Nigeria {
	return &Nigeria{}
}

// KYC returns methods for KYB (Know Your Business)
func (c *Client) KYB() *Kyb {
	return &Kyb{}
}

/*
newRequest makes a http request to youverify server and decodes the server response into the resp(esponse) parameter passed into the newRequest method
*/
func newRequest(method, reqURL string, reqBody, resp interface{}) error {
	newURL := baseURL + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "http client ::: unable to marshal request struct")
		}

		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", secretToken)

	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	defer res.Body.Close()

	if err := parseStatusCode(res.StatusCode); err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	return nil
}

// ParseStatusCode checks the status code response from the server and returns an error for that status code accordingly
func parseStatusCode(code int) error {
	if code == 400 {
		return StatusCode(ErrStatusCode400)
	} else if code == 401 {
		return StatusCode(ErrStatusCode401)
	} else if code == 402 {
		return StatusCode(ErrStatusCode402)
	} else if code == 404 {
		return StatusCode(ErrStatusCode404)
	} else if code == 405 {
		return StatusCode(ErrStatusCode405)
	} else if code == 408 {
		return StatusCode(ErrStatusCode408)
	} else if code == 424 {
		return StatusCode(ErrStatusCode424)
	} else if code == 429 {
		return StatusCode(ErrStatusCode429)
	} else if code == 500 {
		return StatusCode(ErrStatusCode500)
	} else if code == 502 {
		return StatusCode(ErrStatusCode502)
	} else if code == 503 {
		return StatusCode(ErrStatusCode503)
	} else if code == 504 {
		return StatusCode(ErrStatusCode504)
	}

	return nil
}
