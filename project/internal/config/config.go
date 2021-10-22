package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

func New(envFile string) (*Config, error) {

	_ = godotenv.Load(envFile)

	config := Config{}
	if err := envdecode.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
