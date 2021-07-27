package config

import (
	"os"
)

//config data

type Config struct {
	Protocol string
	Host     string
	Port     string
}

func Set() *Config {
	var config Config
	config.Protocol = os.Getenv("PROTOCOL")
	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")

	if config.Protocol == "" {
		config.Protocol = "tcp"
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = ":8080"
	}

	return &Config{
		Protocol: config.Protocol,
		Host:     config.Host,
		Port:     config.Port,
	}
}
