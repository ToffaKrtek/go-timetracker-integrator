package report

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ToffaKrtek/go-timetracker-integrator/internal/service"
)

type TimeReport struct {
  UsageTime map[string]time.Duration `json:"usage_time"`
  FullTime  time.Duration `json:"full_time"`
}

func (tr *TimeReport) Merge(other *TimeReport) {
  (*tr).FullTime += other.FullTime
  for appName, t := range other.UsageTime {
    if _, exists := tr.UsageTime[appName]; exists {
      (*tr).UsageTime[appName] += t
    }else {
      (*tr).UsageTime[appName] = t
    }
  }
}

func GetTimeReport(filename string) (*TimeReport, error)  {
  var tr TimeReport
  file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tr)
	if err != nil {
		return nil, err
	}

	return &tr, nil
}

func createTimeReport(filename string, tr *TimeReport) error {
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()
  encoder := json.NewEncoder(file)
  err = encoder.Encode(tr)
  return err
} 

func SaveTimeReport(
  filename string,
  fulltime time.Duration, 
  usageTime map[string]time.Duration,
) error {
   newTr := TimeReport{
     UsageTime: usageTime,
     FullTime: fulltime,
   }
  tr, err := GetTimeReport(filename)
  if err != nil {
    return createTimeReport(filename, &newTr)
  }
  tr.Merge(&newTr)
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer file.Close()
  encoder := json.NewEncoder(file)
  err = encoder.Encode(tr)
  return err
}

type YoutrackReport struct {
  Closed []service.Task
  Opened []service.Task
  InWork []service.Task
}


func SendReport() {

}

