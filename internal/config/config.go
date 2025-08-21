package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	StoragePath string `yaml:"storagePath"`
	GRPC
}

// todo поменять путь

type GRPC struct {
	Port    string `yaml:"port"`
	Timeout string `yaml:"timeout"`
	Secret  string `yaml:"secret"`
}

func MustLoadConfig() *Config {
	configPath := "config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file doesn't exist at path: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
