package crlayer

import "os"

type Config struct {
	Url    string
	ApiKey string
}

func NewConfigFromEnv() *Config {
	return &Config{Url: os.Getenv("CRLAYER_URL"), ApiKey: os.Getenv("CRLAYER_KEY")}
}
