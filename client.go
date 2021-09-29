package datadoglogsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

const content = "application/json"
const maxSize = 2*1024*1024 - 51
const maxMessageSize = 256 * 1024

type DataDogHttpClientConfiguration struct {
	apiKey string
	host   string
}

type DatadogClient interface {
	Send(entry *logrus.Entry) error
}

type DatadogHttpClient struct {
	formatter  DataDogLogFormater
	datadogUrl string
}

func NewDatadogHttpClient(config DataDogHttpClientConfiguration, formatter DataDogLogFormater) DatadogHttpClient {
	return DatadogHttpClient{formatter: formatter, datadogUrl: fmt.Sprintf("https://%s/api/v1/input/%s", config.host, config.apiKey)}
}

func (c *DatadogHttpClient) Send(entry *logrus.Entry) error {
	msg, err := c.formatter.Format(entry)
	if err != nil {
		return err
	}
	json_data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(c.datadogUrl, content, bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("DataDog Http Error, Status Code: %d", resp.StatusCode)
	}
	return nil
}
