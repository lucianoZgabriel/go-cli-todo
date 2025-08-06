package main

import (
	"encoding/json"
	"fmt"
	"os"
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

// SaveToFile salva a TodoList em um arquivo JSON
func (tl *TodoList) SaveToFile(filename string) error {
	// Converter struct para JSON
	data, err := json.MarshalIndent(tl, "", " ")
	if err != nil {
		return fmt.Errorf("erro ao converter para JSON: %w", err)
	}

	//Escrever arquivo
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever arquivo '%s': %w", filename, err)
	}

	fmt.Printf("💾 TodoList salva em '%s' com sucesso!\n", filename)
	return nil
}

// LoadFromFile carrega TodoList de um arquivo JSON
func (tl *TodoList) LoadFromFile(filename string) error {
	//Verificar se arquivo existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("arquivo '%s' não encontrado", filename)
	}

	// Ler arquivo
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo '%s': %w", filename, err)
	}

	// Converter JSON para struct
	err = json.Unmarshal(data, tl)
	if err != nil {
		return fmt.Errorf("erro ao converter JSON: %w", err)
	}

	fmt.Printf("📁 TodoList carregada de '%s' com sucesso!\n", filename)
	fmt.Printf("📊 %d tarefa(s) carregada(s)\n", len(tl.Tasks))
	return nil
}

// ListPendingTasks exibe apenas tarefas pendentes
func (tl *TodoList) ListPendingTasks() {
	pendingTasks := []Task{}

	for _, task := range tl.Tasks {
		if !task.Completed {
			pendingTasks = append(pendingTasks, task)
		}
	}

	if len(pendingTasks) == 0 {
		fmt.Println("🎉 Nenhuma tarefa pendente! Tudo concluído!")
		return
	}

	fmt.Printf("📋 Tarefas pendentes: %d\n", len(pendingTasks))
	fmt.Println(strings.Repeat("-", 50))

	for _, task := range pendingTasks {
		fmt.Printf("❌ [%d] %s\n", task.ID, task.Title)
		if task.Description != "" {
			fmt.Printf("    📄 %s\n", task.Description)
		}
		fmt.Printf("    🕒 Criada em: %s\n", task.CreatedAt.Format("02/01/2006 15:04"))
		fmt.Println()
	}
}

// ListCompletedTasks exibe apenas tarefas concluídas
func (tl *TodoList) ListCompletedTasks() {
	completedTasks := []Task{}

	for _, task := range tl.Tasks {
		if task.Completed {
			completedTasks = append(completedTasks, task)
		}
	}

	if len(completedTasks) == 0 {
		fmt.Println("📝 Nenhuma tarefa concluída ainda.")
		return
	}

	fmt.Printf("📋 Tarefas concluídas: %d\n", len(completedTasks))
	fmt.Println(strings.Repeat("-", 50))

	for _, task := range completedTasks {
		fmt.Printf("✅ [%d] %s\n", task.ID, task.Title)
		if task.Description != "" {
			fmt.Printf("    📄 %s\n", task.Description)
		}
		fmt.Printf("    🕒 Criada em: %s\n", task.CreatedAt.Format("02/01/2006 15:04"))
		fmt.Println()
	}
}

// SearchTasks busca tarefas por texto no título ou descrição
func (tl *TodoList) SearchTasks(query string) {
	if query == "" {
		fmt.Println("❌ Digite um termo para buscar!")
		return
	}

	query = strings.ToLower(query) //case-insensitive search
	foundTasks := []Task{}

	for _, task := range tl.Tasks {
		//Buscar no título
		titleMatch := strings.Contains(strings.ToLower(task.Title), query)
		//Buscar na descrição
		descMatch := strings.Contains(strings.ToLower(task.Description), query)

		if titleMatch || descMatch {
			foundTasks = append(foundTasks, task)
		}
	}

	if len(foundTasks) == 0 {
		fmt.Printf("🔍 Nenhuma tarefa encontrada para '%s'\n", query)
		return
	}

	fmt.Printf("🔍 Encontradas %d tarefa(s) para '%s':\n", len(foundTasks), query)
	fmt.Println(strings.Repeat("-", 50))

	for _, task := range foundTasks {
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

// ShowStatistics exibe estatísticas das tarefas
func (tl *TodoList) ShowStatistics() {
	total := len(tl.Tasks)
	if total == 0 {
		fmt.Println("📊 Nenhuma tarefa cadastrada ainda.")
		return
	}

	completed := 0
	pending := 0

	for _, task := range tl.Tasks {
		if task.Completed {
			completed++
		} else {
			pending++
		}
	}

	// Calcular percentuais
	completedPercent := float64(completed) / float64(total) * 100
	pendingPercent := float64(pending) / float64(total) * 100

	fmt.Println("📊 === ESTATÍSTICAS DAS TAREFAS ===")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("📝 Total de tarefas:     %d\n", total)
	fmt.Printf("✅ Tarefas concluídas:   %d (%.1f%%)\n", completed, completedPercent)
	fmt.Printf("❌ Tarefas pendentes:    %d (%.1f%%)\n", pending, pendingPercent)
	fmt.Println(strings.Repeat("=", 40))

	// Status geral
	switch completed {
	case total:
		fmt.Println("🎉 Parabéns! Todas as tarefas foram concluídas!")
	case 0:
		fmt.Println("💪 Vamos começar! Marque suas primeiras tarefas como concluídas.")
	default:
		fmt.Printf("🚀 Continue assim! Faltam apenas %d tarefa(s) para concluir tudo.\n", pending)
	}
}

// EditTask permite editar título e descrição de uma tarefa
func (tl *TodoList) EditTask(id int, newTitle string, newDescription string) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			//Validar se pelo menos o título foi fornecido
			if newTitle == "" {
				return fmt.Errorf("❌ Título não pode ser vazio")
			}
			oldTitle := tl.Tasks[i].Title
			oldDesc := tl.Tasks[i].Description

			tl.Tasks[i].Title = newTitle
			tl.Tasks[i].Description = newDescription

			fmt.Printf("✏️ Tarefa [%d] editada com sucesso!\n", id)
			fmt.Printf("   Título: '%s' → '%s'\n", oldTitle, newTitle)
			if oldDesc != newDescription {
				if oldDesc == "" {
					fmt.Printf("   Descrição adicionada: '%s'\n", newDescription)
				} else if newDescription == "" {
					fmt.Printf("   Descrição removida: '%s'\n", oldDesc)
				} else {
					fmt.Printf("   Descrição: '%s' → '%s'\n", oldDesc, newDescription)
				}
			}
			return nil
		}
	}
	return fmt.Errorf("❌ Tarefa com ID %d não encontrada", id)
}
