package bot

import (
	"gopkg.in/telebot.v4"
	"log"
	commands "parser/internal/bot"
	"time"
)

func Start() error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  "6645421211:AAGj5GUzTn3YsiPcrpYj_nd4DzhmzczF6NE",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	commands.BotCommands(bot)
	return err
}
