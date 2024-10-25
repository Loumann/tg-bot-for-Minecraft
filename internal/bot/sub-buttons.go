package commands

import "gopkg.in/telebot.v4"

func createInlineKeyboard() [][]telebot.InlineButton {
	killButton := telebot.InlineButton{
		Unique: "kill",
		Text:   "Убить игрока!",
	}
	giveButton := telebot.InlineButton{
		Unique: "give_item",
		Text:   "Дать игроку палку",
	}
	gameModeButton := telebot.InlineButton{
		Unique: "creative",
		Text:   "Креатив!",
	}

	return [][]telebot.InlineButton{
		{gameModeButton},
		{giveButton},
		{killButton},
	}
}

func ButtonsCommandForWorld() [][]telebot.InlineButton {
	setTime := telebot.InlineButton{
		Unique: "set_time",
		Text:   "Установить время",
	}
	return [][]telebot.InlineButton{
		{setTime},
	}

}
