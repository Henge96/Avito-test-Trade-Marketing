package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Server  Server  `yaml:"server"`
		DB      DB      `yaml:"db"`
		Timeout Timeout `yaml:"timeout"`
	}
	Server struct {
		Host    string `yaml:"host"`
		Port    Port   `yaml:"port"`
		Timeout Timeout
	}
	Port struct {
		Http string `yaml:"http"`
	}
	Timeout struct {
		ServerTimeout int `yaml:"server_timeout"`
	}
	DB struct {
		Driver   string `yaml:"driver"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		HostDB   string `yaml:"hostdb"`
		PortDB   string `yaml:"portdb"`
		NameDB   string `yaml:"namedb"`
		SSLMode  string `yaml:"SSLMode"`
	}
)

func TakeConfigFromYaml(s string) (*Config, error) {
	file, err := os.Open(s)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	var config = &Config{}

	yaml.NewDecoder(file).Decode(config)

	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}

	if config.Server.Port.Http == "" {
		config.Server.Port.Http = "8080"
	}

	if config.Server.Timeout.ServerTimeout == 0 {
		config.Server.Timeout.ServerTimeout = 30
	}

	if config.DB.Driver == "" {
		config.DB.Driver = "postgres"
	}

	if config.DB.User == "" {
		config.DB.User = "postgres"
	}

	if config.DB.Password == "" {
		config.DB.Password = "postgres"
	}

	if config.DB.HostDB == "" {
		config.DB.HostDB = "localhost"
	}

	if config.DB.PortDB == "" {
		config.DB.PortDB = "5432"
	}

	if config.DB.NameDB == "" {
		config.DB.NameDB = "postgres"
	}

	if config.DB.SSLMode == "" {
		config.DB.SSLMode = "disable"
	}

	return config, nil
}
