package task

import (
	"fmt"
	"strings"
	"time"
)

// Task representa uma tarefa individual
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// String implementa a interface Stringer para formatação
func (t *Task) String() string {
	status := "❌ Pendente"
	if t.Completed {
		status = "✅ Concluída"
	}
	return fmt.Sprintf("[%d] %s - %s (%s)",
		t.ID, t.Title, t.Description, status)
}

// TodoList gerencia uma coleção de tasks
type TodoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

// NewTodoList cria uma nova lista de tarefas
func NewTodoList() *TodoList {
	return &TodoList{
		Tasks:  make([]Task, 0),
		NextID: 1,
	}
}

// AddTask adiciona uma nova tarefa à lista
func (tl *TodoList) AddTask(title, description string) *Task {
	task := Task{
		ID:          tl.NextID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++

	return &task
}

// ToggleTask alterna o status de uma tarefa
func (tl *TodoList) ToggleTask(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Completed = !tl.Tasks[i].Completed
			return nil
		}
	}
	return fmt.Errorf("tarefa com ID %d não encontrada", id)
}

// RemoveTask remove uma tarefa da lista
func (tl *TodoList) RemoveTask(id int) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tarefa com ID %d não encontrada", id)
}

// GetTask retorna uma tarefa por ID
func (tl *TodoList) GetTask(id int) (*Task, error) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			return &tl.Tasks[i], nil
		}
	}
	return nil, fmt.Errorf("tarefa com ID %d não encontrada", id)
}

// ListPendingTasks retorna apenas tarefas pendentes
func (tl *TodoList) ListPendingTasks() []Task {
	var pending []Task
	for _, task := range tl.Tasks {
		if !task.Completed {
			pending = append(pending, task)
		}
	}
	return pending
}

// SearchTasks busca tarefas que contenham o termo no título ou descrição
func (tl *TodoList) SearchTasks(query string) []Task {
	var results []Task
	query = strings.ToLower(query)

	for _, task := range tl.Tasks {
		title := strings.ToLower(task.Title)
		description := strings.ToLower(task.Description)

		if strings.Contains(title, query) || strings.Contains(description, query) {
			results = append(results, task)
		}
	}
	return results
}

// Stats retorna estatísticas da lista
func (tl *TodoList) Stats() (total, completed, pending int) {
	total = len(tl.Tasks)
	for _, task := range tl.Tasks {
		if task.Completed {
			completed++
		}
	}
	pending = total - completed
	return
}
