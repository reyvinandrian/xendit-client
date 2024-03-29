package xenditgo

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// Client struct
type Client struct {
	APIEnvType               EnvironmentType
	SecretAPIKey             string
	ApiVersion               string
	InvoiceDurationInSeconds int

	LogLevel int
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	return Client{
		APIEnvType: Sandbox,

		// LogLevel is the logging level used by the library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		LogLevel: 3,
	}
}

// ===================== HTTP CLIENT ================================================
var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, body io.Reader) (*http.Request, error) {
	logLevel := c.LogLevel
	log := clog.Get()

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			log.Error("Request creation failed ", err)
		}
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.SecretAPIKey, "")
	req.Header.Add("api-version", c.ApiVersion)
	log.Debugf("Request %s : %s %s", req.Method, req.Header.Get("api-version"), req.GetBody)
	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request, v interface{}) (httpStatus int, err error) {
	logLevel := c.LogLevel
	log := clog.Get()
	res, err := httpClient.Do(req)
	resBody, err := ioutil.ReadAll(res.Body)
	if logLevel > 1 {
		log.Debugf("Payment response %s", string(resBody))
		log.Debugf("Request %s : %s %s", req.Method, req.URL.Host, req.URL.Path)
	}

	return res.StatusCode, nil
}

// NewRequest for batch disbursement : send new request
func (c *Client) NewDisbBatchRequest(key string, method string, fullPath string, body io.Reader) (*http.Request, error) {
	logLevel := c.LogLevel
	log := clog.Get()

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			log.Error("Request creation failed ", err)
		}
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-IDEMPOTENCY-KEY", key)
	req.SetBasicAuth(c.SecretAPIKey, "")
	return req, nil
}

// Call the Xendit API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
/*
func (c *Client) Call(method, path string, body io.Reader, v interface{}) (httpStatus int, err error) {
	req, err := c.NewRequest(method, path, body)

	if err != nil {
		return httpStatus, err
	}

	return c.ExecuteRequest(req, v)
}
*/
// ===================== END HTTP CLIENT ================================================
