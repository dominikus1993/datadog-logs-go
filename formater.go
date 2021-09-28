package datadoglogsgo

import (
	"os"

	"github.com/sirupsen/logrus"
)

type DataDogLogMessage struct {
	Ddsource string `json:"ddsource"`
	Ddtags   string `json:"ddtags"`
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
	Service  string `json:"service"`
	Level    string `json:"level"`
}

func NewDataDogLogMessage(log *logrus.Entry) *DataDogLogMessage {
	hostname, _ := os.Hostname()
	return &DataDogLogMessage{
		Ddsource: log.Data["source"].(string),
		Ddtags:   log.Data["ddtags"].(string),
		Hostname: hostname,
		Message:  log.Message,
		Service:  log.Data["service"].(string),
		Level:    log.Level.String(),
	}
}

type DataDogLogFormater interface {
	Format(entry *logrus.Entry) (*DataDogLogMessage, error)
}

type dataDogLogFormater struct {
}

func (f *dataDogLogFormater) Format(entry *logrus.Entry) (*DataDogLogMessage, error) {
	return NewDataDogLogMessage(entry), nil
}

func NewDataDogLogFormater() *dataDogLogFormater {
	return &dataDogLogFormater{}
}
