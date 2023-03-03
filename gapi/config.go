package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppStaticConfig struct {
	Dir    string `json:"dir,omitempty"`
	Url    string `json:"url,omitempty"`
	Browse bool   `json:"browse,omitempty"`
}

type AppConfig struct {
	Addr   string           `yaml:"addr" json:"addr,omitempty"`
	MaxCpu int              `yaml:"max_cpu" json:"max_cpu,omitempty"`
	Static *AppStaticConfig `json:"static,omitempty"`
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
