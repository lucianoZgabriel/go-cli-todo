package main

import (
	"fmt"
	"os"

	"github.com/lucianoZgabriel/go-cli-todo/internal/cli"
	"github.com/lucianoZgabriel/go-cli-todo/internal/storage"
)

// Nome padrão do arquivo onde as tarefas serão salvas
const defaultFileName = "tasks.json"

func main() {
	// 1. Cria a camada de Storage (implementação JSON)
	jsonStorage := storage.NewJSONStorage(defaultFileName)

	// 2. Cria a CLI injetando o Storage
	todoApp := cli.NewCLI(jsonStorage)

	// 3. Inicia a aplicação
	if err := todoApp.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Erro ao executar aplicação: %v\n", err)
		os.Exit(1)
	}
}
