package main

import (
	"fmt"
)

func main() {
	fmt.Println("🚀 Todo CLI - Teste de Persistência JSON")
	fmt.Println("Versão: 0.3.0 - Etapa 3: Persistência de Dados")
	fmt.Println()

	filename := "todos.json"

	// PARTE 1: Criar dados e salvar
	fmt.Println("=== PARTE 1: Criando e Salvando Dados ===")
	todoList1 := NewTodoList()

	todoList1.AddTask("Estudar Go", "Aprender conceitos básicos da linguagem Go")
	todoList1.AddTask("Implementar CLI", "Criar interface de linha de comando")
	todoList1.AddTask("Escrever testes", "Criar testes unitários")

	// Marcar uma como concluída
	todoList1.ToggleTask(2)

	fmt.Println()
	fmt.Println("Lista original:")
	todoList1.ListTasks()
	fmt.Println()

	// Salvar no arquivo
	err := todoList1.SaveToFile(filename)
	if err != nil {
		fmt.Printf("Erro ao salvar: %v\n", err)
		return
	}
	fmt.Println()

	// PARTE 2: Carregar dados em nova instância
	fmt.Println("=== PARTE 2: Carregando Dados em Nova Instância ===")
	todoList2 := NewTodoList()

	fmt.Println("Lista vazia antes de carregar:")
	todoList2.ListTasks()
	fmt.Println()

	// Carregar do arquivo
	err = todoList2.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("Erro ao carregar: %v\n", err)
		return
	}
	fmt.Println()

	fmt.Println("Lista após carregar do arquivo:")
	todoList2.ListTasks()
	fmt.Println()

	// PARTE 3: Testar se dados são idênticos
	fmt.Println("=== PARTE 3: Validação de Consistência ===")
	fmt.Printf("Lista original: %d tarefas\n", len(todoList1.Tasks))
	fmt.Printf("Lista carregada: %d tarefas\n", len(todoList2.Tasks))
	fmt.Printf("NextID original: %d\n", todoList1.NextID)
	fmt.Printf("NextID carregado: %d\n", todoList2.NextID)

	// Testar operação na lista carregada
	fmt.Println()
	fmt.Println("Testando operação na lista carregada...")
	todoList2.AddTask("Nova tarefa", "Adicionada após carregar")
	todoList2.ListTasks()

	// PARTE 4: Testar erro de arquivo inexistente
	fmt.Println()
	fmt.Println("=== PARTE 4: Teste de Error Handling ===")
	todoListError := NewTodoList()
	err = todoListError.LoadFromFile("arquivo_inexistente.json")
	if err != nil {
		fmt.Printf("Erro esperado: %v\n", err)
	}
}
