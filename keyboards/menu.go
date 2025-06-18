package keyboards

import "gopkg.in/telebot.v4"

const (
	BtnAppointmentText = "Записаться"
	BtnProfileText     = "Профиль"
)

func MainMenu() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

	// Создаем кнопки через меню
	btnAppointment := menu.Text(BtnAppointmentText)
	btnProfile := menu.Text(BtnProfileText)

	menu.Reply(
		menu.Row(btnAppointment),
		menu.Row(btnProfile),
	)

	return menu
}
