package report

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeReportMerge(t *testing.T) {
  usageTime := make(map[string]time.Duration)
  usageTime["app-1"] = time.Second
  usageTime["app-2"] = time.Second
  tr := TimeReport{
    FullTime: time.Second,
    UsageTime: usageTime,
  }
  usageTime2 := make(map[string]time.Duration)
  usageTime2["app-1"] = time.Second
  usageTime2["app-3"] = time.Second
  tr2 := TimeReport{
    FullTime: time.Second * 2,
    UsageTime: usageTime2,
  }
  tr.Merge(&tr2)
  assert := assert.New(t)
  usageTimeRes := make(map[string]time.Duration)
  usageTimeRes["app-1"] = time.Second * 2
  usageTimeRes["app-2"] = time.Second
  usageTimeRes["app-3"] = time.Second
  assert.Equal(TimeReport{
    FullTime: time.Second * 3,
    UsageTime:  usageTimeRes,
  }, tr)
}

func TestGetTimeReport(t *testing.T) {
  file, err := os.Create("test-time-report-*.json")
  if err != nil {
    t.Errorf("Ошибка создания файла для тестирования: %s", err)
    return
  }
  filename := file.Name()
  defer os.Remove(filename)
  defer file.Close()

  usageTime := make(map[string]time.Duration)
  usageTime["app-1"] = time.Second
  expected := TimeReport{
    FullTime: time.Second,
    UsageTime: usageTime,
  }
  encoder := json.NewEncoder(file)
  if err := encoder.Encode(expected); err != nil {
    t.Errorf("Ошибка записи в файл для тестирования: %s", err)
    return
  }

  tr, err := GetTimeReport(filename)
  if err != nil {
    t.Errorf("Ошибка получения данных из GetTimeReport: %s", err)
    return
  }
  assert := assert.New(t)
  assert.Equal(expected, *tr)
}

func TestCreateTimeReport(t *testing.T) {
  usageTime := make(map[string]time.Duration)
  usageTime["app-1"] = time.Second
  tr := TimeReport{
    FullTime: time.Second,
    UsageTime: usageTime,
  }
  filename := "test-time-report-*.json"
  defer os.Remove(filename)
  err := createTimeReport(filename, &tr)
  assert := assert.New(t)
  assert.Nil(err)
}

func TestSendReport(t *testing.T) {

}

func TestSaveTimeReport(t *testing.T) {
  //TODO:: какой-то тест
}
