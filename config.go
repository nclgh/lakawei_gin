package lakawei_gin

import "os"

var (
	ServiceName string
	ServiceAddr string
	ServicePort string
)

const (
	HttpServiceName = "SERVICE_NAME"
	HttpServiceAddr = "SERVICE_ADDR"
	HttpServicePort = "SERVICE_PORT"
)

func initConfigFromENV() {
	if v := os.Getenv(HttpServiceName); v != "" {
		ServiceName = v
	}

	if v := os.Getenv(HttpServiceAddr); v != "" {
		ServiceAddr = v
	}

	if v := os.Getenv(HttpServicePort); v != "" {
		ServicePort = v
	}
}

func initConfigFromFile() {
	c := GetConfig()

	if v := c.DefaultString("ServiceName", ""); v != "" {
		ServiceName = v
	}

	if v := c.DefaultString("ServiceAddr", ""); v != "" {
		ServiceAddr = v
	}

	if v := c.DefaultString("ServicePort", ""); v != "" {
		ServicePort = v
	}
}

func checkConfig() {
	if ServiceName == "" || ServiceAddr == "" || ServicePort == "" {
		panic("gin config load miss")
	}
}
