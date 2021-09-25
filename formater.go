package datadoglogsgo

import "github.com/sirupsen/logrus"

type DataDogLogMessage struct {
}

type DataDogLogFormater interface {
	Format(entry *logrus.Entry) (*DataDogLogMessage, error)
}
