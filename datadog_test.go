package datadoglogsgo

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type FakeDataDogClient struct {
	message *logrus.Entry
}

func (f *FakeDataDogClient) Send(message *logrus.Entry) error {
	f.message = message
	return nil
}

func NewFakeDataDogClient() *FakeDataDogClient {
	return &FakeDataDogClient{}
}

func TestLocalhostAddAndPrint(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	client := NewFakeDataDogClient()
	cfg := NewDatadogConfiguration("xD", "", []string{"env:dev"})
	hook := NewDatadogHook(cfg, client)

	log.Hooks.Add(hook)

	log.Info("test")

	fmt.Println(client.message.Message)

	assert.Equal(t, "test", client.message.Message)
	assert.Equal(t, "go", client.message.Data["source"])
	assert.Equal(t, "env:dev", client.message.Data["ddtags"])
	assert.Equal(t, "xD", client.message.Data["service"])
	log.Info("Congratulations!")
}
