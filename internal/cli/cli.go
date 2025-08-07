package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lucianoZgabriel/go-cli-todo/internal/storage"
	"github.com/lucianoZgabriel/go-cli-todo/internal/task"
)

// CLI representa a interface de linha de comando
type CLI struct {
	todoList *task.TodoList
	storage  storage.Storage
	scanner  *bufio.Scanner
}

// NewCLI cria uma nova instância da CLI
func NewCLI(storage storage.Storage) *CLI {
	return &CLI{
		todoList: task.NewTodoList(),
		storage:  storage,
		scanner:  bufio.NewScanner(os.Stdin),
	}
}

// Start inicia o loop principal da aplicação
func (c *CLI) Start() error {
	// Carrega dados do storage
	if err := c.loadData(); err != nil {
		return fmt.Errorf("erro ao carregar dados: %w", err)
	}

	fmt.Println("=== 📋 Todo CLI ===")
	fmt.Println("Bem-vindo ao seu gerenciador de tarefas!")

	for {
		c.displayMenu()
		choice := c.readInput("Escolha uma opção: ")

		if err := c.handleMenuChoice(choice); err != nil {
			if err.Error() == "exit" {
				break
			}
		}
	}

	fmt.Println("👋 Até mais!")

	return nil
}

// displayMenu mostra o menu principal
func (c *CLI) displayMenu() {
	total, completed, pending := c.todoList.Stats()

	fmt.Printf("\n=== MENU PRINCIPAL ===\n")
	fmt.Printf("📊 Status: %d total | ✅ %d concluídas | ⏳ %d pendentes\n\n",
		total, completed, pending)

	fmt.Println("1. 📝 Adicionar tarefa")
	fmt.Println("2. 📋 Listar todas as tarefas")
	fmt.Println("3. ✅ Marcar tarefa como concluída")
	fmt.Println("4. ❌ Marcar tarefa como pendente")
	fmt.Println("5. 🗑️  Remover tarefa")
	fmt.Println("6. 🔍 Buscar tarefas")
	fmt.Println("7. ⏳ Listar tarefas pendentes")
	fmt.Println("8. 💾 Salvar e sair")
	fmt.Printf("\n")
}

// handleMenuChoice processa a escolha do usuário
func (c *CLI) handleMenuChoice(choice string) error {
	var err error

	switch choice {
	case "1":
		err = c.addTask()
	case "2":
		err = c.listAllTasks()
	case "3":
		err = c.toggleTaskCompleted(true)
	case "4":
		err = c.toggleTaskCompleted(false)
	case "5":
		err = c.removeTask()
	case "6":
		err = c.searchTasks()
	case "7":
		err = c.listPendingTasks()
	case "8":
		if err := c.saveData(); err != nil {
			return fmt.Errorf("erro ao salvar: %w", err)
		}
		fmt.Println("💾 Dados salvos com sucesso!")
		return fmt.Errorf("exit") // Signal to exit
	default:
		return fmt.Errorf("opção inválida: %s", choice)
	}

	// Se houve erro, mostra e pausa
	if err != nil {
		fmt.Printf("❌ Erro: %s\n", err)
	}

	// SEMPRE pausa após qualquer ação (exceto sair)
	c.waitForEnter()
	return nil
}

// readInput lê uma linha de input do usuário
func (c *CLI) readInput(prompt string) string {
	fmt.Print(prompt)
	if c.scanner.Scan() {
		return strings.TrimSpace(c.scanner.Text())
	}
	return ""
}

// readInt lê um número inteiro do usuário
func (c *CLI) readInt(prompt string) (int, error) {
	input := c.readInput(prompt)
	if input == "" {
		return 0, fmt.Errorf("entrada vazia")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("número inválido: %s", input)
	}

	return num, nil
}

// waitForEnter pausa até o usuário pressionar Enter
func (c *CLI) waitForEnter() {
	fmt.Print("\n🔄 Pressione Enter para continuar...")
	c.scanner.Scan()
}

// loadData carrega dados do storage
func (c *CLI) loadData() error {
	todoList, err := c.storage.Load()
	if err != nil {
		return err
	}

	c.todoList = todoList
	return nil
}

// saveData salva dados no storage
func (c *CLI) saveData() error {
	return c.storage.Save(c.todoList)
}
