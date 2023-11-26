package main

import (
	"fmt"
	"log"
	"strconv"
	"telegram-bot/menu"
	"telegram-bot/tconfig"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	config, err := tconfig.ParseConfig("etc/config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	pref := telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	tele, err := telebot.NewBot(pref)

	// tele.Use(middleware.Logger())
	if err != nil {
		log.Fatal(err)
		return
	}

	menu.LoadAllMenu(tele)

	tele.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send(strconv.FormatInt(c.Chat().ID, 10) + c.Chat().Username)
	})

	// tags
	tele.Handle("/tags", func(c telebot.Context) error {
		tags := c.Args()
		args := ""
		for _, tag := range tags {
			args += tag
		}
		return c.Send(args)
	})

	// payload
	tele.Handle("/clean", func(c telebot.Context) error {
		fmt.Println()
		return nil
	})

	tele.Start()
}
