package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	// TELEGRAM
	TelegramToken   string `json:"telegramToken"`
	TelegramChatId  string `json:"telegramChatId"`
	TelegramTopicId string `json:"telegramTopicId"`

	// YOUTRACK
	YoutrackToken string `json:"youtrackToken"`
	YoutrackUrl   string `json:"youtrackUrl"`

	//User
	UserName    string  `json:"userName"`
	UserWatcher Watcher `json:"userWatcher"`
}

type Watcher int

const (
	HyprlandWatcher Watcher = iota
	XorgWatcher
	WaylandWatcher
	WindowsWatcher
)

var configFileName string = "go-timetracker-integrator.conf"

func (c Config) String() string {
	return fmt.Sprintf(
		"UserName: %s, TgToken: %s, TgChatId: %s, TgTopicId: %s, YTtoken: %s, YTUrl: %s",
		c.UserName, c.TelegramToken, c.TelegramChatId, c.TelegramTopicId, c.YoutrackToken, c.YoutrackUrl,
	)
}

func GetConfig(writer io.Writer, reader *bufio.Reader) *Config {
	var config Config
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		saveConfig(createConfig(writer, reader))
	}
	configFile, err := os.Open(configFileName)
	if err != nil {
		panic("Файл конфигурации недоступен")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		panic("Не удалось прочитать конфигурацию1")
	}
	return &config
}

func saveConfig(config Config) {
	fmt.Println(config)
	file, err := os.Create(configFileName)
	if err != nil {
		panic("Не удалось создать файл конфигурации")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		panic("Не удалось сохранить конфигурацию")
	}
}

func createConfig(writer io.Writer, reader *bufio.Reader) Config {

	var (
		useTelegramIntegration string
		telegramToken          string
		telegramChatId         string
		telegramTopicId        string
		useYoutrackIntegration string
		youtrackToken          string
		youtrackUrl            string
		userName               string
	)

	userName = inputFromCli(writer, reader, "Введите отображаемое имя: ")

	useTelegramIntegration = inputFromCli(writer, reader, "Использовать отправку отчетов в Телеграм (y - да): ")

	if useTelegramIntegration == "y" {

		telegramToken = inputFromCli(writer, reader, "Токен телеграм: ")
		telegramChatId = inputFromCli(writer, reader, "ID чата телеграм:")
		telegramTopicId = inputFromCli(writer, reader, "ID топика: ")
	}

	useYoutrackIntegration = inputFromCli(writer, reader, "Использовать интеграцию Youtrack (y - да): ")
	if useYoutrackIntegration == "y" {
		youtrackToken = inputFromCli(writer, reader, "Токен пользователя Youtrack:")
		youtrackUrl = inputFromCli(writer, reader, "Url вашего youtrack: ")
	}
	return Config{
		UserName:        userName,
		TelegramToken:   telegramToken,
		TelegramChatId:  telegramChatId,
		TelegramTopicId: telegramTopicId,
		YoutrackToken:   youtrackToken,
		YoutrackUrl:     youtrackUrl,
	}
}

func inputFromCli(writer io.Writer, reader *bufio.Reader, q string) string {
	writer.Write([]byte(q))
	value, _ := reader.ReadString('\n')
	return value[:len(value)-1]
}
