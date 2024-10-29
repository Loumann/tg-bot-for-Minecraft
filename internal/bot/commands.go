package commands

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"parser/internal/service"
	"strings"
)

var awaitingInputType = make(map[int64]string)
var players = make(map[int64]string)
var currentPlayerName string

const (
	AwaitingPlayerName       = "player_name"
	AwaitingOtherMessage     = "message"
	AwaitingMessageForPlayer = "message_for_player"
)

func BotCommands(bot *telebot.Bot) {

	ButtonsCommandForPlayer := ButtonsCommandForPlayer1()
	ButtonsCommandForWorld := ButtonsCommandForWorld()
	ButtonsCommandAdditional := ButtonsCommandAdditional()
	ButtonsCommandGamemode := ButtonsCommandGamemode()
	ButtonSelectSetTime := ButtonSelectSetTime()

	bot.Handle("/SetName", func(c telebot.Context) error {
		awaitingInputType[c.Chat().ID] = AwaitingPlayerName
		return c.Send("Введите имя игрока: ")
	})

	bot.Handle("/player", func(ctx telebot.Context) error {
		return ctx.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandForPlayer})
	})
	bot.Handle(&telebot.InlineButton{Unique: "kill"}, func(c telebot.Context) error {
		return HandePlayers(c, "kill "+currentPlayerName)
	})
	bot.Handle(&telebot.InlineButton{Unique: "give_item"}, func(c telebot.Context) error {
		command := fmt.Sprintf("give %s stick 3030", currentPlayerName)
		return HandePlayers(c, command)
	})
	bot.Handle(&telebot.InlineButton{Unique: "gamemode"}, func(c telebot.Context) error {
		return c.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandGamemode})

	})

	bot.Handle("/additionalcommands", func(c telebot.Context) error {
		return c.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandAdditional})
	})
	bot.Handle(&telebot.InlineButton{Unique: "say_message_all_player"}, func(c telebot.Context) error {
		awaitingInputType[c.Chat().ID] = AwaitingOtherMessage
		c.Send("Введите сообщение:")
		return HandePlayers(c, "")
	})
	bot.Handle(&telebot.InlineButton{Unique: "tell_message_for_player"}, func(c telebot.Context) error {
		awaitingInputType[c.Chat().ID] = AwaitingMessageForPlayer
		c.Send("Введите сообщение в следующем формате: playerName message")
		err := HandePlayers(c, "")
		if err != nil {
			fmt.Println(err)
			return err
		}
	})

	bot.Handle("/world", func(c telebot.Context) error {
		return c.Send("Выберите действие:", &telebot.ReplyMarkup{InlineKeyboard: ButtonsCommandForWorld})
	})
	bot.Handle(&telebot.InlineButton{Unique: "set_time"}, func(c telebot.Context) error {
		return c.Edit("Выберите время суток: ", &telebot.ReplyMarkup{InlineKeyboard: ButtonSelectSetTime})

	})

	bot.Handle(telebot.OnCallback, func(c telebot.Context) error {
		unique := strings.TrimSpace(c.Callback().Data)
		fmt.Println("Received unique data:", unique)
		switch unique {

		//Команды для установки времени стуок.
		case "0", "6000", "12000", "18000":
			return HandePlayers(c, "time set "+unique)

		//Команды для управления сервером.
		case "restart", "stop":
			return HandePlayers(c, unique)

		//Команды для режима игры игрока
		case "survival", "creative", "spectator":
			return HandePlayers(c, unique)

		default:
			return c.Respond(&telebot.CallbackResponse{Text: "Неизвестное действие"})
		}
	})

	bot.Handle("/playerlist", func(c telebot.Context) error {
		resp := service.MinecraftCommands(c, "list")
		return c.Send(resp)
	})
	bot.Handle("/setname", func(c telebot.Context) error {
		awaitingInputType[c.Chat().ID] = AwaitingPlayerName // Устанавливаем флаг для ожидания имени игрока
		return c.Send("Введите имя игрока:")
	})
	bot.Handle("/spawn", func(c telebot.Context) error {
		resp := RetrievePlayerLocation(c)
		return c.Send(resp)

	})
	bot.Handle(telebot.OnText, func(c telebot.Context) error {

		chatID := c.Chat().ID

		inputType, awaiting := awaitingInputType[chatID]
		if !awaiting {
			c.Send("Бот ничего не ожидает")
		}
		switch inputType {
		case AwaitingPlayerName:
			playerName := c.Message().Text
			players[chatID] = playerName
			currentPlayerName = playerName
			awaitingInputType[chatID] = ""
			return c.Send(fmt.Sprintf("Имя игрока сохранено: %s", currentPlayerName))

		case AwaitingOtherMessage:
			messageToBroadcast := c.Message().Text

			awaitingInputType[chatID] = ""
			command := fmt.Sprintf("say %s", messageToBroadcast)
			return HandePlayers(c, command)
			c.Send(fmt.Sprintf("Cообщение [%s] отправлено.", messageToBroadcast))

		case AwaitingMessageForPlayer:
			messageForPlaer := c.Message().Text

			awaitingInputType[chatID] = ""
			command := fmt.Sprintf("tell %s", messageForPlaer)
			return HandePlayers(c, command)
			c.Send(fmt.Sprintf("Сообщение для игрока отправлено %s", messageForPlaer))

		default:
			return c.Send("Неизвестный тип ввода.")
		}

		return nil
	})

	bot.Handle(&telebot.InlineButton{Unique: ""}, func(c telebot.Context) error {
		return HandePlayers(c, "")
	})

	bot.Start()
}

func HandePlayers(c telebot.Context, command string) error {
	if currentPlayerName == "" {
		c.Send("Установите имя игрока /SetName")
		return nil
	}
	result := service.MinecraftCommands(c, command)
	fmt.Println(result)
	return c.Respond(&telebot.CallbackResponse{Text: "Ok"})
}

func RetrievePlayerLocation(c telebot.Context) error {
	err := service.MinecraftCommands(c, currentPlayerName)
	return err
}
