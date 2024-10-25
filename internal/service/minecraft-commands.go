package service

import (
	"github.com/gorcon/rcon"
	"gopkg.in/telebot.v4"
	"log"
	"parser/internal/models"
)

func MinecraftCommands(c telebot.Context, command string) error {

	env := models.ParseEvn()

	conn, err := rcon.Dial(env.RconnAddress, env.RconnPassword)
	if err != nil {
		log.Fatalf("Ошибка подключения к RCON: %v", err)
	}

	resp, err := conn.Execute(command)

	if err != nil {
		c.Send("не получилось", err.Error())
		return err
	}
	c.Send(resp)
	defer conn.Close()
	return err

}
