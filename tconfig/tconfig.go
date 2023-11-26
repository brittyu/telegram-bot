package tconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Tconfig struct {
	Name string `json:"name"`
}

func ParseConfig(configName string) (config Tconfig, err error) {
	dataBytes, err := os.ReadFile(configName)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		return
	}

	return
}
