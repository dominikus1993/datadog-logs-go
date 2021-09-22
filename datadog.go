package datadoglogsgo

import (
	"log"
)

type DatadogLogWriter struct {
	apiKey  string
	source  string
	service string
	host    string
	tags    []string
	useSSL  bool
	useTCP  bool
	port    int
}

func NewDatadogLogWriter() *DatadogLogWriter {
	return &DatadogLogWriter{}
}

func (datadog *DatadogLogWriter) Write(p []byte) (n int, err error) {
	log.Print(string(p))
	return len(p), nil
}
