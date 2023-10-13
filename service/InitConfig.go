package service

import (
	"fmt"
	"ginLearnDemo/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig() *model.Config {
	co := &model.Config{}
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Error init Config: %v", err))
	}
	err = yaml.Unmarshal(configFile, co)
	if err != nil {
		panic(fmt.Sprintf("Error init Config: %v", err))
	}
	return co
}
