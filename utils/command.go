package utils

import (
	"fmt"
	"strconv"

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
		var respText string

		for _, ele := range res.Result {
			respText += fmt.Sprintf("/selectline %s %s %s -- %s\n", ele.ID, ele.LineName, ele.StartStationName, ele.EndStationName)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, respText)
		return &msg
	case "selectline":
		models.User.SetAction("selectline")
		line := message.CommandArguments()
		lineID, err := strconv.Atoi(line)
		var text string
		if err != nil {
			text = "🌝 Param error, Integer is required!"
		} else {
			models.User.SelectLine(lineID)
			text = "🎉 Selected!"
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		return &msg
	case "stations":
		line := message.CommandArguments()
		lineID, err := strconv.Atoi(line)
		var text string
		if err != nil {
			lineID = models.User.SelectedLine()
		} else {
			models.User.SelectLine(lineID)
		}
		res, err := LineDetail(lineID)
		if err != nil {
			text = fmt.Sprintf("🌝 %v", err)
		} else {
			text += fmt.Sprintf("🚍 %s\n %s -- %s\n", res.Result.LineName, res.Result.StartStationName, res.Result.EndStationName)
			for _, ele := range res.Result.Stations {
				text += fmt.Sprintf("📍 id: %s %s\n", ele.ID, ele.StationName)
			}
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		return &msg
	case "buses":
		//
		line := message.CommandArguments()
		stations := map[int]models.Station{}
		var text string
		lineID, err := strconv.Atoi(line)
		if err != nil {
			lineID = models.User.SelectedLine()
		} else {
			models.User.SelectLine(lineID)
		}

		if !models.CachedLine.IsExists(lineID) {
			res, err := LineDetail(lineID)
			if err != nil {
				text = fmt.Sprintf("🌝 %v", err)
				msg := tgbotapi.NewMessage(message.Chat.ID, text)
				return &msg
			}
			for _, ele := range res.Result.Stations {
				id, _ := strconv.Atoi(ele.ID)
				stations[id] = ele
			}
			models.CachedLine.Push(lineID, &stations)
		}

		busResp, err := BusDetail(lineID)
		//
		if err != nil || busResp.Status.Code != 0 {
			if err != nil {
				text = fmt.Sprintf("🌝 %v", err)
			} else {
				text = busResp.Status.Msg
			}
			msg := tgbotapi.NewMessage(message.Chat.ID, text)
			return &msg
		}
		//
		cachedStations, _ := models.CachedLine.Get(lineID)
		for _, ele := range busResp.Result {
			text += fmt.Sprintf("🚍 %s %s\n", ele.BusID, (*cachedStations)[ele.StationSeqNum].StationName)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		return &msg

	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Default")
		return &msg
	}
}
