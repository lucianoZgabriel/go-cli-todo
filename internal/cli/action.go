package cli

import (
	"fmt"
	"strings"

	"github.com/lucianoZgabriel/go-cli-todo/internal/task"
)

// addTask adiciona uma nova tarefa
func (c *CLI) addTask() error {
	fmt.Println("\n=== 📝 ADICIONAR NOVA TAREFA ===")

	title := c.readInput("📌 Título da tarefa: ")
	if title == "" {
		return fmt.Errorf("título não pode ser vazio")
	}

	description := c.readInput("📄 Descrição da tarefa: ")
	if description == "" {
		return fmt.Errorf("descrição não pode ser vazia")
	}

	task := c.todoList.AddTask(title, description)
	fmt.Printf("\n✅ Tarefa criada com sucesso!\n")
	fmt.Printf("🆔 ID: %d\n", task.ID)
	fmt.Printf("📌 Título: %s\n", task.Title)
	fmt.Printf("📄 Descrição: %s\n", task.Description)

	return nil
}

// listAllTasks lista todas as tarefas
func (c *CLI) listAllTasks() error {
	fmt.Println("\n=== 📋 TODAS AS TAREFAS ===")

	if len(c.todoList.Tasks) == 0 {
		fmt.Println("📭 Nenhuma tarefa encontrada!")
		return nil
	}

	fmt.Printf("📊 Total de tarefas: %d\n\n", len(c.todoList.Tasks))

	for _, task := range c.todoList.Tasks {
		c.displayTask(&task)
		fmt.Println() // Linha em branco entre tarefas
	}

	return nil
}

// listPendingTasks lista apenas tarefas pendentes
func (c *CLI) listPendingTasks() error {
	fmt.Println("\n=== ⏳ TAREFAS PENDENTES ===")

	pendingTasks := c.todoList.ListPendingTasks()

	if len(pendingTasks) == 0 {
		fmt.Println("🎉 Parabéns! Todas as tarefas foram concluídas!")
		return nil
	}

	fmt.Printf("⏳ Tarefas pendentes: %d\n\n", len(pendingTasks))

	for _, task := range pendingTasks {
		c.displayTask(&task)
		fmt.Println()
	}

	return nil
}

// toggleTaskCompleted alterna o status de uma tarefa
func (c *CLI) toggleTaskCompleted(markAsCompleted bool) error {
	status := "concluída"
	emoji := "✅"
	if !markAsCompleted {
		status = "pendente"
		emoji = "⏳"
	}

	fmt.Printf("\n=== %s MARCAR TAREFA COMO %s ===\n", emoji, strings.ToUpper(status))

	// Primeiro, mostra as tarefas disponíveis
	if len(c.todoList.Tasks) == 0 {
		fmt.Println("📭 Nenhuma tarefa encontrada!")
		return nil
	}

	fmt.Println("📋 Tarefas disponíveis:")
	for _, task := range c.todoList.Tasks {
		c.displayTaskSummary(&task)
	}
	fmt.Println()

	id, err := c.readInt("🆔 Digite o ID da tarefa: ")
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	// Verifica se a tarefa existe
	task, err := c.todoList.GetTask(id)
	if err != nil {
		return err
	}

	// Verifica se a mudança é necessária
	if task.Completed == markAsCompleted {
		currentStatus := "pendente"
		if task.Completed {
			currentStatus = "concluída"
		}
		return fmt.Errorf("tarefa já está %s", currentStatus)
	}

	// Alterna o status
	if err := c.todoList.ToggleTask(id); err != nil {
		return err
	}

	fmt.Printf("\n%s Tarefa marcada como %s!\n", emoji, status)
	fmt.Printf("📌 %s\n", task.Title)

	return nil
}

// removeTask remove uma tarefa
func (c *CLI) removeTask() error {
	fmt.Println("\n=== 🗑️ REMOVER TAREFA ===")

	if len(c.todoList.Tasks) == 0 {
		fmt.Println("📭 Nenhuma tarefa encontrada!")
		return nil
	}

	fmt.Println("📋 Tarefas disponíveis:")
	for _, task := range c.todoList.Tasks {
		c.displayTaskSummary(&task)
	}
	fmt.Println()

	id, err := c.readInt("🆔 Digite o ID da tarefa para remover: ")
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	// Verifica se a tarefa existe antes de remover
	task, err := c.todoList.GetTask(id)
	if err != nil {
		return err
	}

	// Confirmação de remoção
	fmt.Printf("\n⚠️  Tem certeza que deseja remover esta tarefa?\n")
	fmt.Printf("📌 %s\n", task.Title)
	fmt.Printf("📄 %s\n", task.Description)

	confirmation := c.readInput("Digite 'sim' para confirmar: ")
	if strings.ToLower(confirmation) != "sim" {
		fmt.Println("❌ Remoção cancelada.")
		return nil
	}

	if err := c.todoList.RemoveTask(id); err != nil {
		return err
	}

	fmt.Printf("🗑️ Tarefa removida com sucesso!\n")
	return nil
}

// searchTasks busca tarefas por termo
func (c *CLI) searchTasks() error {
	fmt.Println("\n=== 🔍 BUSCAR TAREFAS ===")

	if len(c.todoList.Tasks) == 0 {
		fmt.Println("📭 Nenhuma tarefa encontrada!")
		return nil
	}

	query := c.readInput("🔍 Digite o termo de busca: ")
	if query == "" {
		return fmt.Errorf("termo de busca não pode ser vazio")
	}

	results := c.todoList.SearchTasks(query)

	if len(results) == 0 {
		fmt.Printf("❌ Nenhuma tarefa encontrada para '%s'\n", query)
		return nil
	}

	fmt.Printf("✅ Encontradas %d tarefa(s) para '%s':\n\n", len(results), query)

	for _, task := range results {
		c.displayTask(&task)
		fmt.Println()
	}

	return nil
}

// displayTask exibe uma tarefa completa
func (c *CLI) displayTask(t *task.Task) {
	status := "⏳ Pendente"
	if t.Completed {
		status = "✅ Concluída"
	}

	fmt.Printf("🆔 ID: %d\n", t.ID)
	fmt.Printf("📌 Título: %s\n", t.Title)
	fmt.Printf("📄 Descrição: %s\n", t.Description)
	fmt.Printf("📊 Status: %s\n", status)
	fmt.Printf("📅 Criada em: %s\n", t.CreatedAt.Format("02/01/2006 15:04"))
}

// displayTaskSummary exibe um resumo da tarefa
func (c *CLI) displayTaskSummary(t *task.Task) {
	status := "⏳"
	if t.Completed {
		status = "✅"
	}
	fmt.Printf("  %s [%d] %s\n", status, t.ID, t.Title)
}
