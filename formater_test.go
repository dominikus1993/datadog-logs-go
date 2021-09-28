package datadoglogsgo

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type FormattingDataDogClient struct {
	message *dataDogLogMessage
}

func (f *FormattingDataDogClient) Send(message *logrus.Entry) error {
	formatter := NewDataDogLogFormater()
	msg, err := formatter.Format(message)
	f.message = msg
	return err
}

func NewFormattingDataDogClient() *FormattingDataDogClient {
	return &FormattingDataDogClient{}
}

func TestHookWithFormatter(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	client := NewFormattingDataDogClient()
	cfg := NewDatadogConfiguration("xD", "", []string{"env:dev"})
	hook := NewDatadogHook(cfg, client)

	log.Hooks.Add(hook)

	log.WithField("test", "test").Info("test")

	fmt.Println(client.message.Message)

	//assert.Equal(t, "{\"level\":\"info\",\"msg\":\"test\",\"test\":\"test\",\"time\":\"2021-09-28T16:19:30Z\"}\n", client.message.Message)
	assert.Equal(t, "go", client.message.Ddsource)
	assert.Equal(t, "env:dev", client.message.Ddtags)
	assert.Equal(t, "xD", client.message.Service)
	assert.Equal(t, "info", client.message.Level)
	log.Info("Congratulations!")
}
