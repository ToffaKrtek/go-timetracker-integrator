package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	configFileName = "test-config-*.conf"
	defer os.Remove(configFileName)

	input := "John Doe\n" + // Имя пользователя
		"y\n" + // Использовать Telegram
		"my_telegram_token\n" + // Токен Telegram
		"123456789\n" + // ID чата Telegram
		"my_topic_id\n" + // ID топика
		"y\n" + // Использовать Youtrack
		"my_youtrack_token\n" + // Токен Youtrack
		"http://my.youtrack.url\n" // URL Youtrack

	reader := bufio.NewReader(bytes.NewBufferString(input))
	var output bytes.Buffer

	config := GetConfig(&output, reader)

	expected := Config{
		UserName:        "John Doe",
		TelegramToken:   "my_telegram_token",
		TelegramChatId:  "123456789",
		TelegramTopicId: "my_topic_id",
		YoutrackToken:   "my_youtrack_token",
		YoutrackUrl:     "http://my.youtrack.url",
	}

	if *config != expected {
		t.Errorf("expected %+v, got %+v", expected, *config)
	}
}

func TestSaveConfig(t *testing.T) {
	// Создаем временный файл
	tempFile, err := os.CreateTemp("", "test-config-*.json")
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Удаляем временный файл после теста

	// Устанавливаем имя файла конфигурации на временный файл
	configFileName = tempFile.Name()

	// Создаем тестовую конфигурацию
	testConfig := Config{
		UserName:        "John Doe",
		TelegramToken:   "my_telegram_token",
		TelegramChatId:  "123456789",
		TelegramTopicId: "my_topic_id",
		YoutrackToken:   "my_youtrack_token",
		YoutrackUrl:     "http://my.youtrack.url",
	}

	// Сохраняем конфигурацию
	saveConfig(testConfig)

	// Читаем сохраненный файл
	file, err := os.Open(configFileName)
	if err != nil {
		t.Fatalf("Не удалось открыть файл конфигурации: %v", err)
	}
	defer file.Close()

	var savedConfig Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&savedConfig); err != nil {
		t.Fatalf("Не удалось декодировать файл конфигурации: %v", err)
	}

	// Проверяем, что сохраненная конфигурация соответствует ожидаемой
	if savedConfig != testConfig {
		t.Errorf("expected %+v, got %+v", testConfig, savedConfig)
	}
}

func TestCreateConfig(t *testing.T) {
	tests := []struct {
		input    string
		expected Config
	}{
		{
			"John Doe\ny\nmy_telegram_token\n123456789\nmy_topic_id\ny\nmy_youtrack_token\nhttp://my.youtrack.url\n",
			Config{
				UserName:        "John Doe",
				TelegramToken:   "my_telegram_token",
				TelegramChatId:  "123456789",
				TelegramTopicId: "my_topic_id",
				YoutrackToken:   "my_youtrack_token",
				YoutrackUrl:     "http://my.youtrack.url",
			},
		},
		{
			"John Doe\ny\nmy_telegram_token\n123456789\nmy_topic_id\nn\n",
			Config{
				UserName:        "John Doe",
				TelegramToken:   "my_telegram_token",
				TelegramChatId:  "123456789",
				TelegramTopicId: "my_topic_id",
				YoutrackToken:   "",
				YoutrackUrl:     "",
			},
		},
		{
			"John Doe\nn\nn\n",
			Config{
				UserName:        "John Doe",
				TelegramToken:   "",
				TelegramChatId:  "",
				TelegramTopicId: "",
				YoutrackToken:   "",
				YoutrackUrl:     "",
			},
		},
	}

	for _, test := range tests {
		reader := bufio.NewReader(bytes.NewBufferString(test.input))
		var output bytes.Buffer
		config := createConfig(&output, reader)
		if config != test.expected {
			t.Errorf("expected %+v, got %+v", test.expected, config)
		}
	}
}

func TestInputFromCli(t *testing.T) {
	input := "test input\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))
	var output bytes.Buffer
	question := "Enter something: "
	result := inputFromCli(&output, reader, question)
	expected := "test input"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Проверка, что вопрос был выведен правильно
	expectedOutput := "Enter something: "
	if output.String() != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output.String())
	}
}
