package pkg

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	mu    sync.Mutex
	tasks map[string]*Task
}

func NewService() *Service {
	return &Service{
		tasks: make(map[string]*Task),
	}
}

func (s *Service) CreateTask() *Task {
	id := uuid.New().String()
	task := &Task{
		ID:     id,
		Status: StatusPending,
	}

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	go s.runTask(task)

	return task
}

func (s *Service) GetTask(id string) (*Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task, ok := s.tasks[id]
	return task, ok
}

func (s *Service) runTask(task *Task) {
	s.mu.Lock()
	task.Status = StatusRunning
	s.mu.Unlock()

	time.Sleep(3 * time.Minute)

	s.mu.Lock()
	task.Status = StatusCompleted
	task.Result = "Задача завершена"
	s.mu.Unlock()
}
