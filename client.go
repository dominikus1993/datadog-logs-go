package datadoglogsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const content string = "application/json"
const maxSize int = 2*1024*1024 - 51
const maxMessageSize int = 256 * 1024

type addDataDogHeaderTransport struct {
	T             http.RoundTripper
	datadogApiKey string
}

func (adt *addDataDogHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("DD-EVP-ORIGIN", "datadog-logs-go")
	req.Header.Add("DD-EVP-ORIGIN-VERSION", "v1.0.0")
	req.Header.Add("DD-API-KEY", adt.datadogApiKey)
	return adt.T.RoundTrip(req)
}

func newAddHeaderTransport(datadogApiKey string, T http.RoundTripper) *addDataDogHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &addDataDogHeaderTransport{T: T, datadogApiKey: datadogApiKey}
}

func newDataDogClient(datadogApiKey string) *http.Client {
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: newAddHeaderTransport(datadogApiKey, nil),
	}
	return netClient
}

type datadogError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type dataDogHttpClientConfiguration struct {
	apiKey string
	host   string
}

func NewDatadogHttpClientConfiguration(apiKey string, host string) dataDogHttpClientConfiguration {
	return dataDogHttpClientConfiguration{apiKey: apiKey, host: host}
}

func NewApiKeyDatadogHttpClientConfiguration(apiKey string) dataDogHttpClientConfiguration {
	return dataDogHttpClientConfiguration{apiKey: apiKey, host: "http-intake.logs.datadoghq.com"}
}

type DatadogClient interface {
	Send(entry *logrus.Entry) error
}

type datadogHttpClient struct {
	formatter  DataDogLogFormater
	datadogUrl string
	client     *http.Client
}

func newDatadogHttpClient(config dataDogHttpClientConfiguration, formatter DataDogLogFormater) *datadogHttpClient {
	return &datadogHttpClient{formatter: formatter, datadogUrl: fmt.Sprintf("https://%s/v2/input/%s", config.host, config.apiKey), client: newDataDogClient(config.apiKey)}
}

func (c *datadogHttpClient) Send(entry *logrus.Entry) error {
	msg, err := c.formatter.Format(entry)
	if err != nil {
		return err
	}
	json_data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := c.client.Post(c.datadogUrl, content, bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		decoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()
		var data datadogError
		if err := decoder.Decode(&data); err != nil {
			return fmt.Errorf("DataDog Http Error, Status Code: %d", resp.StatusCode)
		}
		return fmt.Errorf("DataDog Http Error, Status Code: %d, Message: %s, Code: %d", resp.StatusCode, data.Message, data.Code)
	}
	return nil
}
