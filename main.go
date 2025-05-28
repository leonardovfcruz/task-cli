package main

import (
	"bufio"
	"fmt"
	"github.com/leonardovfcruz/tasktracker/tasktracker"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := tasktracker.LoadTasks(); err != nil {
		log.Fatal("Erro ao carregar tarefas:", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Adicionar tarefa")
		fmt.Println("2. Listar todas as tarefas")
		fmt.Println("3. Listar tarefas por status")
		fmt.Println("4. Atualizar tarefa")
		fmt.Println("5. Deletar tarefa")
		fmt.Println("6. Marcar tarefa como in-progress")
		fmt.Println("7. Marcar tarefa como done")
		fmt.Println("0. Sair")
		fmt.Print("Escolha uma opção: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Descrição da tarefa: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)
			task := tasktracker.AddTask(desc)
			fmt.Printf("Tarefa adicionada (ID: %d)\n", task.ID)
		case "2":
			tasks := tasktracker.ListTasks()
			if len(tasks) == 0 {
				fmt.Println("Nenhuma tarefa encontrada.")
			}
			for _, t := range tasks {
				printTask(t)
			}
		case "3":
			fmt.Print("Status (todo, in-progress, done): ")
			status, _ := reader.ReadString('\n')
			status = strings.TrimSpace(status)
			tasks := tasktracker.ListTasksByStatus(status)
			if len(tasks) == 0 {
				fmt.Printf("Nenhuma tarefa com status '%s'.\n", status)
			}
			for _, t := range tasks {
				printTask(t)
			}
		case "4":
			fmt.Print("ID da tarefa: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID inválido")
				continue
			}
			fmt.Print("Nova descrição: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)
			if err := tasktracker.UpdateTask(id, desc); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Tarefa atualizada!")
			}
		case "5":
			fmt.Print("ID da tarefa: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID inválido")
				continue
			}
			if err := tasktracker.DeleteTask(id); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Tarefa deletada!")
			}
		case "6":
			fmt.Print("ID da tarefa: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID inválido")
				continue
			}
			if err := tasktracker.MarkInProgress(id); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Tarefa marcada como in-progress!")
			}
		case "7":
			fmt.Print("ID da tarefa: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID inválido")
				continue
			}
			if err := tasktracker.MarkDone(id); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Tarefa marcada como done!")
			}
		case "0":
			if err := tasktracker.SaveTasks(); err != nil {
				fmt.Println("Erro ao salvar tarefas:", err)
			}
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

		// Salva tarefas após cada alteração importante
		if input == "1" || input == "4" || input == "5" || input == "6" || input == "7" {
			if err := tasktracker.SaveTasks(); err != nil {
				fmt.Println("Erro ao salvar tarefas:", err)
			}
		}
	}
}

func printTask(t tasktracker.Task) {
	fmt.Printf("ID: %d | %s | Status: %s | Criada: %s | Atualizada: %s\n",
		t.ID, t.Description, t.Status,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
		t.UpdatedAt.Format("2006-01-02 15:04:05"))
}
