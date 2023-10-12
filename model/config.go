package model

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Config struct {
		BaseUrl string `yaml:"base_url"`
		Port    string `yaml:"port"`
	} `yaml:"config"`
}

func (co *Config) ReadConfig() (*Config, error) {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, co)
	if err != nil {
		return nil, err
	}
	return co, nil
}

func NewConfig() *Config {
	return &Config{}
}
