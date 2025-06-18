package handlers

import (
	"evolve-or-die/keyboards"
	"gopkg.in/telebot.v4"
)

func HandleStart(c telebot.Context) error {
	return c.Send("Выберите действие:", keyboards.MainMenu())
}

func HandleAppointment(c telebot.Context) error {
	return c.Send("Вы нажали команду Записаться", &telebot.ReplyMarkup{RemoveKeyboard: true})
}

func HandleProfile(c telebot.Context) error {
	return c.Send("Вы нажали команду Профиль", &telebot.ReplyMarkup{RemoveKeyboard: true})
}
