# 📋 Go CLI Todo - Gerenciador de Tarefas em Linha de Comando

Um aplicativo de gerenciamento de tarefas desenvolvido em **Go** com foco em **arquitetura limpa**, **boas práticas** e **padrões de design profissionais**.

![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)

## 🎯 **Sobre o Projeto**

Este projeto foi desenvolvido como um **exercício prático de aprendizado** para dominar conceitos fundamentais do Go e aplicar **padrões arquiteturais** utilizados em aplicações profissionais de grande escala.

### **Principais Objetivos de Aprendizado:**
- ✅ **Clean Architecture** com separação clara de responsabilidades
- ✅ **Dependency Injection** e **Inversion of Control**
- ✅ **Interface-driven design** para alta testabilidade
- ✅ **Error handling** idiomático do Go
- ✅ **Package organization** seguindo convenções da comunidade
- ✅ **JSON persistence** com marshaling/unmarshaling

## 🏗️ **Arquitetura**

O projeto utiliza uma **arquitetura em camadas (Layered Architecture)** inspirada nos princípios da Clean Architecture:

```
📁 go-cli-todo/
├── main.go                 # 🎯 Dependency Injection Container
├── 📁 internal/
│   ├── 📁 task/            # 🧠 Domain Layer (Business Logic)
│   │   └── task.go         #    → Task, TodoList, core business rules
│   ├── 📁 storage/         # 💾 Persistence Layer  
│   │   ├── storage.go      #    → Storage interface definition
│   │   └── json.go         #    → JSON implementation
│   └── 📁 cli/             # 🖥️  Presentation Layer
│       ├── cli.go          #    → CLI interface and main loop
│       └── actions.go      #    → User interaction handlers
└── go.mod
```

### **Fluxo de Dados:**
```
User Input → CLI Layer → Domain Layer → Storage Layer → File System
```

### **Benefícios Arquiteturais:**
- 🔄 **Baixo Acoplamento**: Cada camada é independente
- 🧪 **Alta Testabilidade**: Interfaces permitem mocking fácil
- 🔧 **Manutenibilidade**: Mudanças isoladas por responsabilidade
- 📈 **Escalabilidade**: Fácil adição de novos storages (SQL, NoSQL, etc.)
- 👥 **Team Development**: Times podem trabalhar em paralelo nas camadas

## 🚀 **Funcionalidades**

### **Operações CRUD Completas:**
- ➕ **Adicionar** tarefas com título e descrição
- 📋 **Listar** todas as tarefas ou apenas pendentes
- ✅ **Marcar** tarefas como concluídas/pendentes
- 🗑️ **Remover** tarefas com confirmação de segurança
- 🔍 **Buscar** tarefas por termo (título ou descrição)

### **Características Técnicas:**
- 💾 **Persistência JSON** automática
- 📊 **Estatísticas** em tempo real (total, concluídas, pendentes)
- 🎨 **Interface rica** com emojis e formatação
- ⚠️ **Validação robusta** de entrada do usuário
- 🛡️ **Error handling** com mensagens claras

## 🛠️ **Tecnologias e Padrões**

### **Stack Técnica:**
- **Language:** Go 1.19+
- **Architecture:** Clean Architecture / Layered Architecture
- **Patterns:** Dependency Injection, Repository Pattern, Strategy Pattern
- **Storage:** JSON file-based persistence
- **CLI:** Native Go standard library

### **Padrões de Design Aplicados:**
- 🏗️ **Dependency Injection Container** (main.go)
- 🔌 **Repository Pattern** (Storage interface)
- 🎯 **Strategy Pattern** (Multiple storage implementations)
- 🏭 **Factory Pattern** (NewTodoList, NewCLI constructors)
- 📋 **Command Pattern** (CLI action handlers)

## 📦 **Instalação e Execução**

### **Pré-requisitos:**
- Go 1.19 ou superior instalado
- Git para clone do repositório

### **Passos:**

1. **Clone o repositório:**
```bash
git clone https://github.com/lucianoZgabriel/go-cli-todo.git
cd go-cli-todo
```

2. **Execute a aplicação:**
```bash
go run main.go
```

3. **Ou compile e execute:**
```bash
go build -o todo-cli
./todo-cli
```

## 🎮 **Como Usar**

### **Interface do Menu:**
```
=== MENU PRINCIPAL ===
📊 Status: 5 total | ✅ 2 concluídas | ⏳ 3 pendentes

1. 📝 Adicionar tarefa
2. 📋 Listar todas as tarefas  
3. ✅ Marcar tarefa como concluída
4. ❌ Marcar tarefa como pendente
5. 🗑️ Remover tarefa
6. 🔍 Buscar tarefas
7. ⏳ Listar tarefas pendentes
8. 💾 Salvar e sair
```

### **Exemplo de Uso:**
```bash
# Adicionar uma nova tarefa
Escolha uma opção: 1
📌 Título da tarefa: Estudar Clean Architecture
📄 Descrição da tarefa: Revisar conceitos de DDD e SOLID

✅ Tarefa criada com sucesso!
🆔 ID: 1
📌 Título: Estudar Clean Architecture
📄 Descrição: Revisar conceitos de DDD e SOLID
```

## 🧪 **Testabilidade**

O projeto foi arquitetado com **testabilidade como prioridade**:

### **Interfaces para Mock:**
```go
// Fácil de mockar em testes
type Storage interface {
    Save(todoList *task.TodoList) error
    Load() (*task.TodoList, error)
}
```

### **Dependency Injection:**
```go
// Testável via injeção de dependência
cli := NewCLI(mockStorage)  // Mock para testes
cli := NewCLI(jsonStorage)  // Real para produção
```

## 📈 **Evolução do Projeto**

Este projeto evoluiu através de **6 etapas estruturadas**:

1. **Estruturas Básicas** → Definição de Task e TodoList
2. **CRUD em Memória** → Operações básicas sem persistência  
3. **Persistência JSON** → Sistema de armazenamento
4. **Interface CLI Melhorada** → UX e formatação
5. **Recursos Avançados** → Busca, filtros, estatísticas
6. **Refatoração Arquitetural** → Clean Architecture implementada

## 🎓 **Aprendizados Técnicos**

### **Conceitos Go Dominados:**
- 📦 **Package organization** com `internal/`
- 🔗 **Interface composition** e polimorfismo
- ⚡ **Error handling** idiomático (`error` interface)
- 🏗️ **Struct embedding** e method receivers
- 🔄 **JSON marshaling/unmarshaling** com struct tags
- 📝 **Go modules** e dependency management

### **Padrões Arquiteturais:**
- 🎯 **Separation of Concerns** (SoC)
- 🔄 **Dependency Inversion Principle** (DIP)  
- 🔒 **Single Responsibility Principle** (SRP)
- 🔓 **Open/Closed Principle** (OCP)
- 🔧 **Interface Segregation Principle** (ISP)

## 🤝 **Contribuição**

Este projeto está aberto para contribuições! Sinta-se à vontade para:

- 🐛 Reportar bugs
- 💡 Sugerir melhorias
- 🔧 Enviar pull requests
- 📚 Melhorar documentação

## 📄 **Licença**

Este projeto está sob a licença **MIT**. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👤 **Autor**

**Luciano Gabriel**
- 🐙 GitHub: [@lucianoZgabriel](https://github.com/lucianoZgabriel)
- 💼 LinkedIn: [luciano-gabriel](https://linkedin.com/in/luciano-gabriel)

---

⭐ **Se este projeto foi útil para você, considere dar uma estrela!**

*Desenvolvido com 💚 em Go*