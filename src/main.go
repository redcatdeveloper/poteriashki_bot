package main

import (
  "fmt"
  "log"
  "strconv"
  "github.com/redcatdeveloper/poteriashki_bot/constants"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  goconf "github.com/redcatdeveloper/goconf"
)

func main() {
  conf, err := goconf.NewGoConf("config.ini")
  if err != nil {
    log.Panic(err)
  }

  log.Printf("Usage Token: %s \n", conf.Get("Token"))
  DebugChannelId, _ := strconv.ParseInt(conf.Get("DebugChannelId"), 10, 64)

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
			debugLog(bot, DebugChannelId, fmt.Sprintf("[%s] on %d: %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text))

      if update.Message.IsCommand() {
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

        switch update.Message.Command() {
          case "start":
            msg.Text = constants.HELLO_ON_START
          default:
            msg.Text = "ой"
        }
			  bot.Send(msg)
      }
		}
	}
}

func debugLog(bot *tgbotapi.BotAPI, channelId int64, str string) {
  log.Print(str)
  msg := tgbotapi.NewMessage(channelId, str)
  bot.Send(msg)
}
