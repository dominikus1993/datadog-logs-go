package datadoglogsgo

import "github.com/sirupsen/logrus"

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
