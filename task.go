package main

import (
	"time"
)

// Tarefa individual no sistema
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// TodoList gerencia uma coleção de tarefas
type TodoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

// NewTodoList cria uma nova instância de TodoList inicializada
func NewTodoList() *TodoList {
	return &TodoList{
		Tasks:  make([]Task, 0), //Slice vazio, inicializado
		NextID: 1,               //Ids começam em 1
	}
}

// AddTask adiciona uma nova tarefa à lista
func (tl *TodoList) AddTask(title, description string) *Task {
	task := &Task{
		ID:          tl.NextID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tl.Tasks = append(tl.Tasks, *task)
	tl.NextID++

	return task
}
