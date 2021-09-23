package datadoglogsgo

import (
	"log"

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

func (hook *DatadogLogWriter) prepareMessage(entry *logrus.Entry) (string, error) {
	entry.Data["source"] = hook.config.getSource()
	entry.Data["ddtags"] = hook.config.getDDTags()
	entry.Data["service"] = hook.config.service
	entry.Data["host"] = hook.config.host
	return entry.String()
}

func (hook *DatadogLogWriter) Fire(entry *logrus.Entry) error {
	entry.Data["source"] = hook.config.getSource()
	entry.Data["ddtags"] = hook.config.getDDTags()
	entry.Data["service"] = hook.config.service
	entry.Data["host"] = hook.config.host
	entry.String()
	return nil
}
