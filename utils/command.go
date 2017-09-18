package utils

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sliverwing/BusBot/models"
)

func CommandHandler(message *tgbotapi.Message) *tgbotapi.MessageConfig {
	switch message.Command() {
	case "reset":
		models.User.Clear()
		msg := tgbotapi.NewMessage(message.Chat.ID, "Reset Succeed")
		return &msg
	case "search":
		models.User.SetAction("search")
		line := message.CommandArguments()
		res, err := DailSearchLine(line)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(err))
			return &msg
		}
		if res.Status.Code != 0 {
			msg := tgbotapi.NewMessage(message.Chat.ID, res.Status.Msg)
			return &msg
		}
		var respText string

		for _, ele := range res.Result {
			respText += fmt.Sprintf("/selectline %s %s %s -- %s\n", ele.ID, ele.LineName, ele.StartStationName, ele.EndStationName)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, respText)
		return &msg
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Default")
		return &msg
	}
}
