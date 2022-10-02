package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}
}

func InitializeConfig() (*Config, error) {
	conf := &Config{}

	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
