package datadoglogsgo

func NewDatadogHook(config *datadogConfiguration, clientcfg dataDogHttpClientConfiguration) *DatadogHook {
	formatter := newDataDogLogFormater()
	client := newDatadogHttpClient(clientcfg, formatter)
	return newDatadogHook(config, client)
}
