package menu

import "gopkg.in/telebot.v3"

func LoadAllMenu(tele *telebot.Bot) {
	initDefine()

	tele.Handle("/start", start())
}
