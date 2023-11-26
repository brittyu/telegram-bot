package main

import (
	"fmt"
	"telegram-bot/tconfig"
)

var config tconfig.Tconfig
var err error

func main() {
	config, err = tconfig.ParseConfig("etc/config.yaml")
	if err != nil {
		fmt.Println("parse config error:", err)
		return
	}

	fmt.Println("name:", config.Name)
}
