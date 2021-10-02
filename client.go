package datadoglogsgo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const content = "application/json"
const maxSize = 2*1024*1024 - 51
const maxMessageSize = 256 * 1024

var netClient = &http.Client{
	Timeout: time.Second * 10,
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
}

func newDatadogHttpClient(config dataDogHttpClientConfiguration, formatter DataDogLogFormater) *datadogHttpClient {
	return &datadogHttpClient{formatter: formatter, datadogUrl: fmt.Sprintf("https://%s/v1/input/%s", config.host, config.apiKey)}
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
	var buf bytes.Buffer
	g := gzip.NewWriter(&buf)
	if _, err = g.Write(json_data); err != nil {
		return err
	}
	if err = g.Close(); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.datadogUrl, &buf)
	if err != nil {
		return err
	}
	resp, err := netClient.Do(req)
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
