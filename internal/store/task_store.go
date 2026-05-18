package store

import (
	"sync"
	"time"
	"todo-api/internal/model"
)

type TaskStore struct {
	tasks  map[int]*model.Task
	nextID int
	mu     sync.RWMutex
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]*model.Task),
		nextID: 1,
	}
}

func (s *TaskStore) Create(title, description string) *model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &model.Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks[task.ID] = task
	s.nextID++
	return task
}

func (s *TaskStore) GetAll() []model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]model.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, *task)
	}
	return tasks
}

func (s *TaskStore) GetByID(id int) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, exists := s.tasks[id]
	return task, exists
}

func (s *TaskStore) Update(id int, req model.UpdateTaskRequest) (*model.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, false
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Done != nil {
		task.Done = *req.Done
	}

	task.UpdatedAt = time.Now()
	return task, true
}

func (s *TaskStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[id]
	if exists {
		delete(s.tasks, id)
		return true
	}
	return false
}
