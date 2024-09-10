package service

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/mock"
)

type TestTgMsg struct {
  text string
  path string
}

func (t TestTgMsg) GetText() string {
  return t.text
}

func (t TestTgMsg) GetFilePath() string {
  return t.path
}

type MockTgBot struct {
  mock.Mock
}

func (m *MockTgBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
  args := m.Called(c)
  return args.Get(0).(tgbotapi.Message), args.Error(1)
}

type TestTgConfig struct {
  telegramToken string
  chatId int64
}

func (t TestTgConfig) GetTelegramToken() string {
  return ""
}

func (t TestTgConfig) GetChatId() int64 {
  return int64(0)
}

func TestSend(t *testing.T) {

}

func TestSendMsg(t *testing.T){

}

func TestSendFile(t *testing.T){

}

