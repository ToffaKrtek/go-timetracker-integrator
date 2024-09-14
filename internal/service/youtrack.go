package service

import (
	"errors"
	"time"
)
 
type Task struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Status string `json:"Status"`
  Assignee string `json:"assignee"`
  CreatedBy string `json:"createdBy"`
  Updated time.Time `json:"updated"`
  //TODO:: Createtor
}

type YoutrackInterface interface {
  GetTasks() ([]Task, error)
}

var (
  ErrYoutrackPermissionDenied = errors.New("Ошибка доступа к Youtrack")
  ErrYoutrackNotFound = errors.New("Неверный url для Youtrack")
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

func (y *YoutrackService) GetTasks() ([]Task, error) {
  //TODO:: make request to YOUTACK by y.url + y.token
  return []Task{}, nil
}

func Init(conf YoutrackConfig) {
  youtrackHub = &YoutrackService{
    url: conf.GetYoutrackUrl(),
    token: conf.GetYoutrackToken(),
  }
}

func saveTasks(tasks []Task, filename string) ([]Task, error) {
  // TODO:: load from today-file (or create), add tasks analyze
  //exist, err := os.Stat()
  return []Task{}, nil
}
