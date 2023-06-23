package main

import (
  "log"
  "github.com/redcatdeveloper/poteriashki_bot/constants"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  goconf "github.com/redcatdeveloper/goconf"
)

func main() {
  conf := goconf.NewGoConf("config.ini")

	bot, err := tgbotapi.NewBotAPI(conf.Get("Token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, constants.HELLO_ON_START)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
