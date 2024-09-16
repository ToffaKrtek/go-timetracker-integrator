package tracker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ToffaKrtek/go-timetracker-integrator/internal/config"
	"github.com/ToffaKrtek/go-timetracker-integrator/internal/report"
)

var usageTime = make(map[string]time.Duration)
var fullTime time.Duration

func Run(conf *config.Config) {
	filename := time.Now().Format("2006-01-02") + ".json"
	go trackUsage(conf.UserWatcher)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				// Можно оставить пустым, если не нужно записывать в файл каждые 10 секунд
				report.SaveTimeReport(filename, fullTime, usageTime)
			}
		}
	}()

	// Ожидание сигнала завершения
	<-sigs
	report.SendReport()

}
func trackUsage(watcher config.Watcher) {
	var getActivityFunc funcUsage

	switch watcher {
	case config.HyprlandWatcher:
		getActivityFunc = getActivityHyprland()
	case config.XorgWatcher:
		getActivityFunc = getActivityXorg()
	case config.WaylandWatcher:
		getActivityFunc = getActivityWayland()
	case config.WindowsWatcher:
		getActivityFunc = getActivityWindows()
	default:
		getActivityFunc = getActivityHyprland()
	}
	for {
    fullTime += time.Second
		activeApp, err := getActivityFunc()
		if err != nil {
			fmt.Println("Ошибка при получении активного приложения:", err)
			return
		}

		usageTime[activeApp] += time.Second

		time.Sleep(1 * time.Second) // Проверяем каждую секунду
	}
}

type funcUsage func() (string, error)

func getActivityHyprland() funcUsage {
	return func() (string, error) {
		cmd := exec.Command("hyprctl", "activewindow")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("ошибка выполнения hyprctl: %v, stderr: %s", err, stderr.String())
		}
		output := out.String()
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			converted := strings.TrimSpace(line)

			if strings.HasPrefix(converted, "class:") {
				// Извлекаем название класса приложения
				return strings.TrimSpace(strings.Split(converted, ":")[1]), nil
			}
			if strings.HasPrefix(converted, "initialClass:") {
				return strings.TrimSpace(strings.Split(converted, ":")[1]), nil
			}
		}

		// Если не нашли класс, выводим полный вывод для отладки
		return "", fmt.Errorf("активное приложение не найдено, вывод hyprctl: %s", output)

	}
}

func getActivityXorg() funcUsage {
	return func() (string, error) {
		return "", fmt.Errorf("активное приложение не найдено, вывод hyprctl: %s", "x")
	}
}

func getActivityWayland() funcUsage {
	return func() (string, error) {
		return "", fmt.Errorf("активное приложение не найдено, вывод hyprctl: %s", "x")
	}
}

func getActivityWindows() funcUsage {
	return func() (string, error) {
		return "", fmt.Errorf("активное приложение не найдено, вывод hyprctl: %s", "x")
	}
}
