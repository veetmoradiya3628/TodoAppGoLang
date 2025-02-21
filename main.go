package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-cli <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  add <task>        - Add a new todo")
		fmt.Println("  list              - List all todos")
		fmt.Println("  complete <id>     - Mark a todo as completed")
		fmt.Println("  delete <id>       - Delete a todo")
		return
	}

	commnd := os.Args[1]
	switch commnd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli add <task>")
			return
		}
		AddTodo(os.Args[2])
	case "list":
		ListTodos()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli complete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id")
			return
		}
		CompleteTodo(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id")
			return
		}
		DeleteTodo(id)
	default:
		fmt.Println("Invalid command")
	}
}
