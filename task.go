package main

import (
	"fmt"
	"strings"
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

// ListTasks exibe todas as tarefas formatadas
func (tl *TodoList) ListTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada!")
		return
	}

	fmt.Printf("📋 Total de tarefas: %d\n", len(tl.Tasks))
	fmt.Println(strings.Repeat("-", 50))

	for _, task := range tl.Tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}

		fmt.Printf("%s [%d] %s\n", status, task.ID, task.Title)
		if task.Description != "" {
			fmt.Printf("    📄 %s\n", task.Description)
		}
		fmt.Printf("    🕒 Criada em: %s\n", task.CreatedAt.Format("02/01/2006 15:04"))
		fmt.Println()
	}
}

// ToggleTask alterna o status de conclusão de uma tarefa
func (tl *TodoList) ToggleTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Completed = !tl.Tasks[i].Completed
			status := "pendente"
			if tl.Tasks[i].Completed {
				status = "concluída"
			}
			fmt.Printf("✅ Tarefa [%d] marcada como %s!\n", id, status)
			return nil
		}
	}
	return fmt.Errorf("❌ Tarefa com ID %d não encontrada", id)
}

// RemoveTask remove uma tarefa da lista pelo ID
func (tl *TodoList) RemoveTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			fmt.Printf("🗑️ Tarefa [%d] removida com sucesso!\n", id)
			return nil
		}
	}
	return fmt.Errorf("❌ Tarefa com ID %d não encontrada", id)
}
