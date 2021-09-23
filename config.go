package datadoglogsgo

import "strings"

const GO = "go"

type DatadogConfiguration struct {
	apiKey  string
	source  *string
	service string
	host    string
	tags    []string
	useSSL  bool
	useTCP  bool
	port    int
}

func NewDatadogConfiguration() *DatadogConfiguration {
	return &DatadogConfiguration{}
}

func (cfg *DatadogConfiguration) getDDTags() string {
	if cfg.tags == nil {
		return ""
	}
	return strings.Join(cfg.tags, ",")
}

func (cfg *DatadogConfiguration) getSource() string {
	if cfg.source == nil {
		return GO
	}
	return *cfg.source
}
