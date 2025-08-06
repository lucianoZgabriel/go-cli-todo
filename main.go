package main

import (
	"fmt"
)

func main() {
	fmt.Println("🚀 Todo CLI - Iniciando desenvolvimento!")
	fmt.Println("Versão: 0.1.0 - Etapa 1: Estrutura Básica")
	fmt.Println()

	//Criar uma nova TodoList
	todoList := NewTodoList()
	fmt.Printf("TodoList inicial: %+v\n", todoList)
	fmt.Printf("Número inicial de tarefas: %d\n", len(todoList.Tasks))
	fmt.Println()

	//Adicionar Task
	fmt.Println("=== Testando AddTask ===")
	task1 := todoList.AddTask("Estudar Go", "Aprender conceitos básicos da linguagem Go")
	fmt.Printf("Tarefa criada: %+v\n", task1)

	task2 := todoList.AddTask("Implementar CLI", "Criar interface de linha de comando")
	fmt.Printf("Tarefa criada: %+v\n", task2)

	fmt.Println()
	fmt.Printf("TodoList após adicionar tarefas: %+v\n", todoList)
	fmt.Printf("Número de tarefas: %d\n", len(todoList.Tasks))
	fmt.Printf("Próximo ID será: %d\n", todoList.NextID)
}
