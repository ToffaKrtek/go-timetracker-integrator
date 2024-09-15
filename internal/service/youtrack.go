package service

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)
 
type Task struct {
	Updated   time.Time `json:"updated"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"Status"`
	Assignee  string    `json:"assignee"`
	CreatedBy string    `json:"createdBy"`
}

type Tasks map[string]Task

func MakeTasks(tasksSlice []Task) Tasks {
   tasks := make(Tasks)
  for _, task := range tasksSlice {
    if len(task.ID) > 0 {
    tasks[task.ID] = task
    }
  }
  return tasks
}

func (t *Tasks) Merge(other *Tasks) {
  for id, task := range *other {
    (*t)[id] = task
  }
}

type YoutrackInterface interface {
  GetTasks() (*Tasks, error)
}

var (
  ErrYoutrackPermissionDenied = errors.New("ошибка доступа к Youtrack")
  ErrYoutrackNotFound = errors.New("неверный url для Youtrack")
  ErrYoutrackDataFileError = errors.New("ошибка доступа к файлу с данными Youtrack")
  ErrYoutrackDataSaveError = errors.New("ошибка сохранения файла с данными Youtrack")
)

var youtrackHub YoutrackInterface

type YoutrackConfig interface {
  GetYoutrackUrl() string
  GetYoutrackToken() string
}

type YoutrackService struct {
  url string
  token string
}

func (y *YoutrackService) GetTasks() (*Tasks, error) {
  //TODO:: make request to YOUTACK by y.url + y.token
  return &Tasks{}, nil
}

func Init(conf YoutrackConfig) {
  youtrackHub = &YoutrackService{
    url: conf.GetYoutrackUrl(),
    token: conf.GetYoutrackToken(),
  }
}

func getYoutrackDataTasks(filename string) (*Tasks, error) {
  _, err := os.Stat(filename)

  if err != nil {
    if  errors.Is(err, os.ErrNotExist) {
        _, err = os.Create(filename)
    }
    if err != nil {
      return nil, err
      // return nil, ErrYoutrackDataFileError
    }
  }
  var data *Tasks
  youtrackFile, err := os.Open(filename)
	if err != nil {
    return nil, ErrYoutrackDataFileError
	}
	defer youtrackFile.Close() 

  decoder := json.NewDecoder(youtrackFile)
  if err := decoder.Decode(&data); err != nil {
      return nil, err
    return nil, ErrYoutrackDataFileError
  }
  return data, nil
}


func saveTasks(tasks *Tasks, filename string) (*Tasks, error) {
  data, err := getYoutrackDataTasks(filename)
  if err != nil {
    return nil, err
  }
  data.Merge(tasks)
  file, err := os.Create(filename)
  if err != nil {
    return nil, ErrYoutrackDataSaveError
  }
  defer file.Close()
  encoder := json.NewEncoder(file)
  if err := encoder.Encode(data); err != nil {
    return nil, ErrYoutrackDataSaveError
  }
  return data, nil
}
