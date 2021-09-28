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
	return &dataDogLogMessage{
		Ddsource: log.Data["source"].(string),
		Ddtags:   log.Data["ddtags"].(string),
		Hostname: hostname,
		Message:  log.Message,
		Service:  log.Data["service"].(string),
		Level:    log.Level.String(),
	}, nil
}

func NewDataDogLogFormater() *dataDogLogFormater {
	return &dataDogLogFormater{}
}
