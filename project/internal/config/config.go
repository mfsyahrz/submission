package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

func New() (*Config, error) {

	_ = godotenv.Load()

	config := Config{}
	if err := envdecode.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
