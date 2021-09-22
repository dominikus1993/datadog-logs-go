package datadoglogsgo

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type DatadogLogWriter struct {
	config *DatadogConfiguration
	client *DatadogHttpClient
}

func NewDatadogLogWriter() *DatadogLogWriter {
	return &DatadogLogWriter{}
}

func (datadog *DatadogLogWriter) Write(p []byte) (n int, err error) {
	log.Print(string(p))
	return len(p), nil
}

func (hook *DatadogLogWriter) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	return nil
}
