package handlers

import (
	"evolve-or-die/keyboards"
	"gopkg.in/telebot.v4"
)

// Показать главное меню
func ShowMainMenu(c telebot.Context) error {
	return c.Send("Выберите действие:", keyboards.MainMenu())
}

// Убрать клавиатуру
func RemoveKeyboard(c telebot.Context, text string) error {
	return c.Send(
		text,
		&telebot.ReplyMarkup{RemoveKeyboard: true},
	)
}
