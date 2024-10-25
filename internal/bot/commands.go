package commands

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"parser/internal/service"
)

var players = make(map[int64]string)
var awaitingName = make(map[int64]bool)
var currentPlayerName string

func BotCommands(bot *telebot.Bot) {

	ButtonsCommandForPlayer := createInlineKeyboard()
	ButtonsCommandForWorld := ButtonsCommandForWorld()

	bot.Handle("/player", func(ctx telebot.Context) error {
		return ctx.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandForPlayer})
	})
	bot.Handle("/world", func(ctx telebot.Context) error {
		return ctx.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandForWorld})
	})
	bot.Handle("/playerList", func(c telebot.Context) error {
		resp := service.MinecraftCommands(c, "list")
		return c.Send(resp)
	})
	bot.Handle("/SetName", func(c telebot.Context) error {
		awaitingName[c.Chat().ID] = true
		return c.Send("Введите имя игрока: ")
	})

	bot.Handle(&telebot.InlineButton{Unique: "kill"}, func(c telebot.Context) error {
		return HandePlayers(c, "kill ")
	})
	bot.Handle(&telebot.InlineButton{Unique: "give_item"}, func(c telebot.Context) error {
		command := fmt.Sprintf("give %s stick 3030", currentPlayerName)
		return HandePlayers(c, command)
	})
	bot.Handle(&telebot.InlineButton{Unique: "creative"}, func(c telebot.Context) error {
		command := fmt.Sprintf("gamemode creative " + currentPlayerName)
		return HandePlayers(c, command)
	})
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		if awaitingName[c.Chat().ID] {
			playerName := c.Message().Text
			players[c.Chat().ID] = playerName
			awaitingName[c.Chat().ID] = false
			currentPlayerName = playerName
			return c.Send(fmt.Sprintf("Имя игрока сохранено: %s", currentPlayerName))
		}

		return nil
	})

	bot.Start()
}

func HandePlayers(c telebot.Context, command string) error {
	if currentPlayerName == "" {
		c.Send("Установите имя игрока /SetName")
		return nil
	}
	return service.MinecraftCommands(c, command)
}
