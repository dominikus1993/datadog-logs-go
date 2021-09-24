package datadoglogsgo

import "strings"

const GO = "go"

type DatadogConfiguration struct {
	source  string
	service string
	tags    []string
}

func NewDatadogConfiguration(service, source string, tags []string) *DatadogConfiguration {
	return &DatadogConfiguration{service: service, source: source, tags: tags}
}

func (cfg *DatadogConfiguration) getDDTags() string {
	if cfg.tags == nil {
		return ""
	}
	return strings.Join(cfg.tags, ",")
}

func (cfg *DatadogConfiguration) getSource() string {
	if cfg.source == "" {
		return GO
	}
	return cfg.source
}
