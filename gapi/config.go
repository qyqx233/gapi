package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	App *AppConfig `yaml:"app"`
}

func parseYaml(path string) *Config {
	var c = new(Config)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, c)
	return c
}
