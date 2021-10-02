package datadoglogsgo

import (
	"os"

	"github.com/sirupsen/logrus"
)

type dataDogLogMessage struct {
	Ddsource string `json:"ddsource"`
	Ddtags   string `json:"ddtags"`
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
	Service  string `json:"service"`
	Level    string `json:"level"`
}

type DataDogLogFormater interface {
	Format(entry *logrus.Entry) (*dataDogLogMessage, error)
}

type dataDogLogFormater struct {
}

func (f *dataDogLogFormater) Format(log *logrus.Entry) (*dataDogLogMessage, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	source := log.Data["source"].(string)
	delete(log.Data, "source")
	service := log.Data["service"].(string)
	delete(log.Data, "service")
	tags := log.Data["ddtags"].(string)
	delete(log.Data, "ddtags")
	msg, err := log.String()
	if err != nil {
		return nil, err
	}
	return &dataDogLogMessage{
		Ddsource: source,
		Ddtags:   tags,
		Hostname: hostname,
		Message:  msg,
		Service:  service,
		Level:    log.Level.String(),
	}, nil
}

func newDataDogLogFormater() *dataDogLogFormater {
	return &dataDogLogFormater{}
}
