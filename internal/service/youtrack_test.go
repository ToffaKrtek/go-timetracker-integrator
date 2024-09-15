package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type YoutrackMock struct {
  mock.Mock
}

func (m *YoutrackMock) GetTasks() (*Tasks, error) {
  args := m.Called()
  return args.Get(0).(*Tasks), args.Error(1)
}

func TestGet(t *testing.T) {

}

func TestMakeTasks(t *testing.T) {
  assert := assert.New(t)
  tasks := MakeTasks([]Task{
        Task{
      	ID:        "test-1",
      	Title:     "test-1 title",
      	Status:    "closed",
      	Assignee:  "",
      	CreatedBy: "",
      },
        Task{
      	ID:        "test-2",
      	Title:     "test-2 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
  })
    if assert.NotNil(tasks) {
      assert.Equal(tasks, Tasks{
        "test-1": Task{
      	ID:        "test-1",
      	Title:     "test-1 title",
      	Status:    "closed",
      	Assignee:  "",
      	CreatedBy: "",
      },
        "test-2": Task{
      	ID:        "test-2",
      	Title:     "test-2 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
    })
    }
}

func TestSaveTasks(t *testing.T){
  assert := assert.New(t)
  tests := []struct{
    tasks Tasks
    expected Tasks
  }{
    {
    tasks: Tasks{
        "test-1": Task{
      	ID:        "test-1",
      	Title:     "test-1 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
    },
    expected: Tasks{
        "test-1": Task{
      	ID:        "test-1",
      	Title:     "test-1 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
    },
    },
    {
    tasks: Tasks{
        "test-2": Task{
      	ID:        "test-2",
      	Title:     "test-2 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
        "test-1": Task{
      	ID:        "test-1",
      	Title:     "",
      	Status:    "closed",
      	Assignee:  "",
      	CreatedBy: "",
      },
    },
    expected: Tasks{
        "test-1": Task{
      	ID:        "test-1",
      	Title:     "test-1 title",
      	Status:    "closed",
      	Assignee:  "",
      	CreatedBy: "",
      },
        "test-2": Task{
      	ID:        "test-2",
      	Title:     "test-2 title",
      	Status:    "open",
      	Assignee:  "",
      	CreatedBy: "",
      },
    },
    },
  }
  testFileName := fmt.Sprintf("test-%s.json", time.Now().Format("2006-01-02T15:04:05.000000"))
  for _, test := range tests {
    res, err := saveTasks(&test.tasks, testFileName)
    assert.Nil(err)
    if assert.NotNil(res) {
      assert.Equal(*res, test.expected)
    }
  }
  os.Remove(testFileName)
}
