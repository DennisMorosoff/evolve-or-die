package main

import (
	"log"
	"os"
	"time"

	"evolve-or-die/handlers"
	"evolve-or-die/keyboards"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {
	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	botToken := os.Getenv("TELE_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELE_BOT_TOKEN environment variable not set")
	}

	settings := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация обработчиков
	bot.Handle("/start", handlers.HandleStart)
	bot.Handle(keyboards.BtnAppointmentText, handlers.HandleAppointment)
	bot.Handle(keyboards.BtnProfileText, handlers.HandleProfile)

	log.Println("Bot started successfully!")
	bot.Start()
}
