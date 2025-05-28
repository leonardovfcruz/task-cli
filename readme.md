
# Task CLI

*Read this in [Portuguese](#task-cli-pt-br)*

A simple command-line task management application built with Go.

## Features

- Create new tasks with automatic ID assignment
- List all tasks
- Filter tasks by status (todo, in-progress, done)
- Update task descriptions
- Delete tasks
- Change task status (mark as in-progress or done)
- Persistent storage using JSON file

## Requirements

- Go 1.24 or higher

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/leonardovfcruz/tasktracker.git
   cd tasktracker
   ```

2. Build the application:
   ```
   go build -o task-cli
   ```

## Usage

Run the application:
```
./task-cli in Linux/macOS:

tasktracker.exe in Windows
```

### Menu Options

The application presents a menu with the following options:

1. **Add task** - Create a new task with a description
2. **List all tasks** - Display all tasks with their details
3. **List tasks by status** - Filter tasks by status (todo, in-progress, done)
4. **Update task** - Change the description of an existing task
5. **Delete task** - Remove a task by ID
6. **Mark task as in-progress** - Change task status to in-progress
7. **Mark task as done** - Change task status to done
0. **Exit** - Save tasks and exit the application

## Data Storage

Tasks are automatically saved to a `tasks.json` file in the application directory after each modification. When the application starts, it loads existing tasks from this file.

## Example Workflow

Here's an example of how you might use Task CLI to manage your tasks:

```
===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 1

Task description: Complete project documentation
Task added (ID: 1)

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 1

Task description: Fix login bug
Task added (ID: 2)

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 2

ID: 1 | Complete project documentation | Status: todo | Created: 2023-05-10 14:30:45 | Updated: 2023-05-10 14:30:45
ID: 2 | Fix login bug | Status: todo | Created: 2023-05-10 14:31:20 | Updated: 2023-05-10 14:31:20

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 6

Task ID: 2
Task marked as in-progress!

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 3

Status (todo, in-progress, done): in-progress
ID: 2 | Fix login bug | Status: in-progress | Created: 2023-05-10 14:31:20 | Updated: 2023-05-10 14:32:05

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 7

Task ID: 2
Task marked as done!

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 2

ID: 1 | Complete project documentation | Status: todo | Created: 2023-05-10 14:30:45 | Updated: 2023-05-10 14:30:45
ID: 2 | Fix login bug | Status: done | Created: 2023-05-10 14:31:20 | Updated: 2023-05-10 14:33:10

===== MENU =====
1. Add task
2. List all tasks
3. List tasks by status
4. Update task
5. Delete task
6. Mark task as in-progress
7. Mark task as done
0. Exit
Choose an option: 0

Exiting...
```

This workflow demonstrates creating tasks, listing them, changing their status, and filtering by status.

---


<a name="task-cli-pt-br"></a>
# Task CLI (Português)

Uma aplicação simples de gerenciamento de tarefas via linha de comando construída com Go.

## Funcionalidades

- Criar novas tarefas com atribuição automática de ID
- Listar todas as tarefas
- Filtrar tarefas por status (todo, in-progress, done)
- Atualizar descrições de tarefas
- Deletar tarefas
- Alterar status de tarefas (marcar como em andamento ou concluída)
- Armazenamento persistente usando arquivo JSON

## Requisitos

- Go 1.24 ou superior

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/leonardovfcruz/tasktracker.git
   cd tasktracker
   ```

2. Compile a aplicação:
   ```
   go build -o task-cli
   ```

## Uso

Execute a aplicação:
```
./task-cli para Linux/macOS

tasktracker.exe para Windows
```

### Opções do Menu

A aplicação apresenta um menu com as seguintes opções:

1. **Adicionar tarefa** - Cria uma nova tarefa com uma descrição
2. **Listar todas as tarefas** - Exibe todas as tarefas com seus detalhes
3. **Listar tarefas por status** - Filtra tarefas por status (todo, in-progress, done)
4. **Atualizar tarefa** - Altera a descrição de uma tarefa existente
5. **Deletar tarefa** - Remove uma tarefa pelo ID
6. **Marcar tarefa como in-progress** - Altera o status da tarefa para em andamento
7. **Marcar tarefa como done** - Altera o status da tarefa para concluída
0. **Sair** - Salva as tarefas e sai da aplicação

## Armazenamento de Dados

As tarefas são automaticamente salvas em um arquivo `tasks.json` no diretório da aplicação após cada modificação. Quando a aplicação inicia, ela carrega as tarefas existentes deste arquivo.

## Exemplo de Fluxo de Trabalho

Aqui está um exemplo de como você pode usar o Task CLI para gerenciar suas tarefas:

```
===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 1

Descrição da tarefa: Completar documentação do projeto
Tarefa adicionada (ID: 1)

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 1

Descrição da tarefa: Corrigir bug de login
Tarefa adicionada (ID: 2)

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 2

ID: 1 | Completar documentação do projeto | Status: todo | Criada: 2023-05-10 14:30:45 | Atualizada: 2023-05-10 14:30:45
ID: 2 | Corrigir bug de login | Status: todo | Criada: 2023-05-10 14:31:20 | Atualizada: 2023-05-10 14:31:20

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 6

ID da tarefa: 2
Tarefa marcada como in-progress!

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 3

Status (todo, in-progress, done): in-progress
ID: 2 | Corrigir bug de login | Status: in-progress | Criada: 2023-05-10 14:31:20 | Atualizada: 2023-05-10 14:32:05

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 7

ID da tarefa: 2
Tarefa marcada como done!

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 2

ID: 1 | Completar documentação do projeto | Status: todo | Criada: 2023-05-10 14:30:45 | Atualizada: 2023-05-10 14:30:45
ID: 2 | Corrigir bug de login | Status: done | Criada: 2023-05-10 14:31:20 | Atualizada: 2023-05-10 14:33:10

===== MENU =====
1. Adicionar tarefa
2. Listar todas as tarefas
3. Listar tarefas por status
4. Atualizar tarefa
5. Deletar tarefa
6. Marcar tarefa como in-progress
7. Marcar tarefa como done
0. Sair
Escolha uma opção: 0

Saindo...
```

Este fluxo de trabalho demonstra como criar tarefas, listá-las, alterar seus status e filtrar por status.

https://roadmap.sh/projects/task-tracker