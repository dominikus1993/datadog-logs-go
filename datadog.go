package datadoglogsgo

import (
	"log"

	"github.com/sirupsen/logrus"
)

type DatadogHook struct {
	config *DatadogConfiguration
	client DatadogClient
}

func NewDatadogLogWriter(config *DatadogConfiguration, client DatadogClient) *DatadogHook {
	return &DatadogHook{config: config, client: client}
}

func (datadog *DatadogHook) Write(p []byte) (n int, err error) {
	log.Print(string(p))
	return len(p), nil
}

func (hook *DatadogHook) prepareMessage(entry *logrus.Entry) {
	entry.Data["source"] = hook.config.getSource()
	entry.Data["ddtags"] = hook.config.getDDTags()
	entry.Data["service"] = hook.config.service
}

func (hook *DatadogHook) Fire(entry *logrus.Entry) error {
	hook.prepareMessage(entry)
	return hook.client.Send(entry)
}

func (hook *DatadogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
