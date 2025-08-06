package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Função para exibir o menu principal
func displayMenu() {
	fmt.Println("\n🚀 Todo CLI - Gerenciador de Tarefas")
	fmt.Println("=======================================")
	fmt.Println("\n📋 Menu Principal:")
	fmt.Println("1. 📝 Listar todas as tarefas")
	fmt.Println("2. ➕ Adicionar nova tarefa")
	fmt.Println("3. ✅ Marcar tarefa como concluída/pendente")
	fmt.Println("4. 🗑️  Remover tarefa")
	fmt.Println("5. 💾 Salvar dados")
	fmt.Println("6. 📁 Carregar dados")
	fmt.Println("0. 🚪 Sair")
	fmt.Print("\nDigite sua opção: ")
}

// Função para ler entrada do usuário
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Função para ler um número inteiro
func readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	input := readInput()
	return strconv.Atoi(input)
}

func main() {
	fmt.Println("🚀 Todo CLI v0.4.0 - Interface Interativa")
	fmt.Println("Carregando aplicação...")

	// Criar TodoList
	todoList := NewTodoList()
	filename := "todos.json"

	// Tentar carregar dados existentes
	err := todoList.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("📋 Iniciando com lista vazia (arquivo não encontrado)\n")
	}

	// Loop principal da aplicação
	for {
		displayMenu()
		option := readInput()

		switch option {
		case "1":
			// Listar tarefas
			fmt.Println("\n📝 === TODAS AS TAREFAS ===")
			todoList.ListTasks()

		case "2":
			// Adicionar tarefa
			fmt.Println("\n➕ === ADICIONAR NOVA TAREFA ===")
			fmt.Print("📝 Digite o título da tarefa: ")
			title := readInput()

			if title == "" {
				fmt.Println("❌ Título não pode ser vazio!")
				break
			}

			fmt.Print("📄 Digite a descrição (ou Enter para pular): ")
			description := readInput()

			task := todoList.AddTask(title, description)
			fmt.Printf("✅ Tarefa [%d] criada com sucesso!\n", task.ID)

		case "3":
			// Toggle tarefa
			fmt.Println("\n✅ === MARCAR TAREFA ===")
			todoList.ListTasks()

			id, err := readInt("🔢 Digite o ID da tarefa: ")
			if err != nil {
				fmt.Println("❌ ID inválido! Digite apenas números.")
				break
			}

			err = todoList.ToggleTask(id)
			if err != nil {
				fmt.Println(err)
			}

		case "4":
			// Remover tarefa
			fmt.Println("\n🗑️ === REMOVER TAREFA ===")
			todoList.ListTasks()

			id, err := readInt("🔢 Digite o ID da tarefa para remover: ")
			if err != nil {
				fmt.Println("❌ ID inválido! Digite apenas números.")
				break
			}

			err = todoList.RemoveTask(id)
			if err != nil {
				fmt.Println(err)
			}

		case "5":
			// Salvar dados
			fmt.Println("\n💾 === SALVAR DADOS ===")
			err := todoList.SaveToFile(filename)
			if err != nil {
				fmt.Printf("❌ Erro ao salvar: %v\n", err)
			}

		case "6":
			// Carregar dados
			fmt.Println("\n📁 === CARREGAR DADOS ===")
			err := todoList.LoadFromFile(filename)
			if err != nil {
				fmt.Printf("❌ Erro ao carregar: %v\n", err)
			}

		case "0":
			// Sair
			fmt.Println("\n🚪 Encerrando aplicação...")
			fmt.Println("💾 Salvando dados automaticamente...")
			err := todoList.SaveToFile(filename)
			if err != nil {
				fmt.Printf("⚠️  Aviso: Erro ao salvar dados: %v\n", err)
			}
			fmt.Println("👋 Obrigado por usar o Todo CLI!")
			return

		default:
			fmt.Printf("❌ Opção '%s' inválida! Digite um número de 0 a 6.\n", option)
		}

		// Pausa para o usuário ler o resultado
		fmt.Println("\nPressione Enter para continuar...")
		readInput()
	}
}
