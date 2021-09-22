package datadoglogsgo

type DatadogConfiguration struct {
	apiKey  string
	source  string
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
