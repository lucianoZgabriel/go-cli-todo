package main

import (
	"fmt"
)

func main() {
	fmt.Println("🚀 Todo CLI - Teste CRUD Completo!")
	fmt.Println("Versão: 0.2.0 - Etapa 2: CRUD Básico")
	fmt.Println()

	// Criar nova TodoList
	todoList := NewTodoList()

	// CREATE: Adicionar tarefas
	fmt.Println("=== CREATE: Adicionando Tarefas ===")
	todoList.AddTask("Estudar Go", "Aprender conceitos básicos da linguagem Go")
	todoList.AddTask("Implementar CLI", "Criar interface de linha de comando")
	todoList.AddTask("Escrever testes", "Criar testes unitários")
	todoList.AddTask("Deploy", "Colocar aplicação em produção")
	fmt.Println()

	// READ: Listar tarefas
	fmt.Println("=== READ: Listando Todas as Tarefas ===")
	todoList.ListTasks()
	fmt.Println()

	// UPDATE: Marcar algumas como concluídas
	fmt.Println("=== UPDATE: Marcando Tarefas como Concluídas ===")
	err := todoList.ToggleTask(1)
	if err != nil {
		fmt.Println(err)
	}

	err = todoList.ToggleTask(3)
	if err != nil {
		fmt.Println(err)
	}

	// Testar ID inexistente
	err = todoList.ToggleTask(99)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// READ: Ver mudanças
	fmt.Println("=== READ: Ver Tarefas Após Toggle ===")
	todoList.ListTasks()
	fmt.Println()

	// DELETE: Remover uma tarefa
	fmt.Println("=== DELETE: Removendo Tarefa ===")
	err = todoList.RemoveTask(2)
	if err != nil {
		fmt.Println(err)
	}

	// Testar remoção de ID inexistente
	err = todoList.RemoveTask(99)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// READ: Estado final
	fmt.Println("=== READ: Estado Final ===")
	todoList.ListTasks()
}
