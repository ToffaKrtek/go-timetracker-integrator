package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type YoutrackMock struct {
  mock.Mock
}

func (m *YoutrackMock) GetTasks() ([]Task, error) {
  args := m.Called()
  return args.Get(0).([]Task), args.Error(1)
}

func TestGet(t *testing.T) {

}

func TestSaveTasks(t *testing.T){
  assert := assert.New(t)
  tests := []struct{
    tasks []Task
    expected []Task
  }{
    {
    tasks: []Task{
      Task{
        ID: "test-1",
        Title: "test-1 title",
        Status: "open",
      },
    },
    expected: []Task{
      Task{
        ID: "test-1",
        Title: "test-1 title",
        Status: "open",
      },
    },
    },
    {
    tasks: []Task{
      Task{
        ID: "test-2",
        Title: "test-2 title",
        Status: "open",
      },
      Task{
        ID: "test-1",
        Status: "closed",
      },
    },
    expected: []Task{
      Task{
        ID: "test-1",
        Title: "test-1 title",
        Status: "closed",
      },
      Task{
        ID: "test-2",
        Title: "test-2 title",
        Status: "open",
      },
    },
    },
  }
  testFileName := fmt.Sprintf("test-%s.json", time.Now().Format("2006-01-02T15:04:05.000000"))
  for _, test := range tests {
    res, err := saveTasks(test.tasks, testFileName)
    assert.Nil(err)
    if assert.NotNil(res) {
      assert.Equal(res, test.expected)
    }
  }
}
