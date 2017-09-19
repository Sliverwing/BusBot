package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sliverwing/BusBot/models"
	"github.com/sliverwing/BusBot/utils"
)

func main() {

	models.User = &models.UserStatus{}

	bot, err := tgbotapi.NewBotAPI("431548070:AAGEgvuPzkqJC-SlV96rUPGpWR4fTtd7XB0")
	if err != nil {
		log.Panicln(err)
	}

	bot.Debug = true

	log.Printf("Authorized Bot Name: %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			bot.Send(utils.CommandHandler(update.Message))
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
