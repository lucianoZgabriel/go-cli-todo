package storage

import (
	"github.com/lucianoZgabriel/go-cli-todo/internal/task"
)

type Storage interface {
	Save(todoList *task.TodoList) error
	Load() (*task.TodoList, error)
}
