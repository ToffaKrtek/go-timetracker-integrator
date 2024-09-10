package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgMsg interface {
  GetText() string
  GetFilePath() string
}

type TgBot interface {
  Send(tgbotapi.Chattable) (tgbotapi.Message, error)
}

var botApi TgBot
var chatId int64

type TgConfig interface {
  GetTelegramToken() string
  GetChatId() int64
}

func Init(conf TgConfig) {
  bot, err := tgbotapi.NewBotAPI(conf.GetTelegramToken())
  if err != nil {
    panic(err)
  }
  botApi = bot
  chatId = conf.GetChatId()
}

func Send(msg TgMsg) error {
  var err error
  err = sendMsg(msg)
  if len(msg.GetFilePath()) > 0 {
    err = sendFile(msg)
  }
  return err
}

func sendMsg(tgMsg TgMsg) error {
  msg := tgbotapi.NewMessage(chatId, tgMsg.GetText())
  _, err := botApi.Send(msg)
  return err
}

func sendFile(tgMsg TgMsg) error {
    inputFile := tgbotapi.FilePath(tgMsg.GetFilePath())
    msg := tgbotapi.NewDocument(chatId, inputFile)
    _, err := botApi.Send(msg)
  return err
}
