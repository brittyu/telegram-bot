package menu

import "gopkg.in/telebot.v3"

func start() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		return c.Send("startMenu", StartMenu)
	}
}
