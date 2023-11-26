package tconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type RedisConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Tconfig struct {
	Name  string    `json:"name"`
	Token string    `json:"token"`
	Redis RedisConf `json:"redis"`
	Mysql MysqlConf `json:"mysql"`
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
