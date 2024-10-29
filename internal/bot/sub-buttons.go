package commands

import "gopkg.in/telebot.v4"

func ButtonsCommandForPlayer1() [][]telebot.InlineButton {
	killButton := telebot.InlineButton{
		Unique: "kill",
		Text:   "Убить игрока!",
	}
	giveButton := telebot.InlineButton{
		Unique: "give_item",
		Text:   "Дать игроку палку",
	}
	gameModeButton := telebot.InlineButton{
		Unique: "gamemode",
		Text:   "Режим игры",
	}
	clearInventory := telebot.InlineButton{
		Unique: "clear",
		Text:   "Очистить инвентарь",
	}
	next := telebot.InlineButton{
		Unique: "next",
		Text:   "Следующая страница",
	}

	return [][]telebot.InlineButton{
		{killButton}, {gameModeButton},
		{giveButton}, {clearInventory},
		{next},
	}
}

func ButtonsCommandForWorld() [][]telebot.InlineButton {
	setTime := telebot.InlineButton{
		Unique: "set_time",
		Text:   "Установить время",
	}
	weather := telebot.InlineButton{
		Unique: "weather",
		Text:   "Установить погоду",
	}
	difficulty := telebot.InlineButton{
		Unique: "difficulty",
		Text:   "Сложность",
	}
	spawnpoint := telebot.InlineButton{
		Unique: "set_time",
		Text:   "Спавн для игрока",
	}

	return [][]telebot.InlineButton{
		{weather}, {setTime},
		{difficulty},
		{spawnpoint},
	}

}
func ButtonsCommandAdditional() [][]telebot.InlineButton {
	restartSerever := telebot.InlineButton{
		Unique: "restart",
		Text:   "Перезагрузить сервер",
	}
	stopServer := telebot.InlineButton{
		Unique: "stop",
		Text:   "Остановаить сервер",
	}
	sayMessageAllPlayer := telebot.InlineButton{
		Unique: "say_message_all_player",
		Text:   "Сообщение для всех игроков",
	}
	tell := telebot.InlineButton{
		Unique: "tell_message_for_player",
		Text:   "Отправить сообщение игроку",
	}
	return [][]telebot.InlineButton{
		{restartSerever, stopServer},
		{sayMessageAllPlayer},
		{tell},
	}
}
func ButtonsCommandGamemode() [][]telebot.InlineButton {
	survival := telebot.InlineButton{
		Unique: "survival",
		Text:   "выживание",
	}
	creative := telebot.InlineButton{
		Unique: "creative",
		Text:   "Креатив",
	}
	spectator := telebot.InlineButton{
		Unique: "spectator",
		Text:   "Наблюдатель",
	}

	return [][]telebot.InlineButton{
		{creative, survival},
		{spectator},
	}

}
func ButtonSelectSetTime() [][]telebot.InlineButton {
	sunrise := telebot.InlineButton{
		Unique: "0",
		Text:   "Рассвет",
	}
	day := telebot.InlineButton{
		Unique: "6000",
		Text:   "День",
	}
	sunset := telebot.InlineButton{
		Unique: "12000",
		Text:   "Закат",
	}
	night := telebot.InlineButton{
		Unique: "18000",
		Text:   "Ночь",
	}

	return [][]telebot.InlineButton{
		{sunrise, day, sunset, night},
	}
}
