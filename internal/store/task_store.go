package store

import(
	"sync"
	"time"
	"todo-api/internal/model"
)

type TaskStore struct{
	tasks map[int]*model.Task
	nextID int
	mu sync.RWMutex
}

