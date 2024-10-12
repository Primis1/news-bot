package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env string `yaml:"env" env-default:"development"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG_PATH", "exc/internal/local.yaml")

	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(config); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(config, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
