package menu

import "gopkg.in/telebot.v3"

var (
	StartMenu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	StartSelector = &telebot.ReplyMarkup{}
	SettingsBtn   = StartMenu.Text("⚙ Settings")
)

func initDefine() {
	StartMenu.Reply(
		StartMenu.Row(SettingsBtn),
	)
}
