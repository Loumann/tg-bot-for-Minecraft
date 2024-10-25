package main

import (
	"parser/api/telegram-api/bot"
)

func main() {

	err := bot.Start()
	if err != nil {
		panic(err)
	}
}
