package storage

import (
	"encoding/json"
	"os"

	"github.com/lucianoZgabriel/go-cli-todo/internal/task"
)

// JSONStorage implementa a interface Storage usando arquivos JSON
type JSONStorage struct {
	filename string
}

// NewJSONStorage cria um novo storage JSON com o arquivo especificado
func NewJSONStorage(filename string) Storage {
	return &JSONStorage{
		filename: filename,
	}
}

// Save persiste a TodoList em arquivo JSON
func (js *JSONStorage) Save(todoList *task.TodoList) error {
	file, err := os.Create(js.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") // Formatação JSON

	return encoder.Encode(todoList)
}

// Load carrega uma TodoList do arquivo JSON
func (js *JSONStorage) Load() (*task.TodoList, error) {
	// Se arquivo não existe, retorna lista vazia
	if _, err := os.Stat(js.filename); os.IsNotExist(err) {
		return task.NewTodoList(), nil
	}

	file, err := os.Open(js.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var todoList task.TodoList
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&todoList); err != nil {
		return nil, err
	}

	return &todoList, nil
}
