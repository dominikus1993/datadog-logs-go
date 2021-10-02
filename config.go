package datadoglogsgo

import "strings"

const GO = "go"

type datadogConfiguration struct {
	source  string
	service string
	tags    []string
}

func NewDatadogConfiguration(service, source string, tags []string) *datadogConfiguration {
	return &datadogConfiguration{service: service, source: source, tags: tags}
}

func (cfg *datadogConfiguration) getDDTags() string {
	if cfg.tags == nil {
		return ""
	}
	return strings.Join(cfg.tags, ",")
}

func (cfg *datadogConfiguration) getSource() string {
	if cfg.source == "" {
		return GO
	}
	return cfg.source
}
