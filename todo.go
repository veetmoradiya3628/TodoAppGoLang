package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const todoFile = "todos.json"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// LodaTodos reads the todos from the JSON file
func LoadTodos() ([]Todo, error) {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(file, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// SaveTodos saves the todos to the JSON file
func SaveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

// Add a new todo
func AddTodo(title string) {
	todos, _ := LoadTodos()
	id := len(todos) + 1
	todos = append(todos, Todo{ID: id, Title: title, Completed: false})
	SaveTodos(todos)
	fmt.Println("✅ Todo added successfully!")
}

// List all todos
func ListTodos() {
	todos, _ := LoadTodos()
	fmt.Println("📝 Your Todos:")
	for _, todo := range todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("[%s] %d %s\n", status, todo.ID, todo.Title)
	}
}

// Mark a todo as completed
func CompleteTodo(id int) {
	todos, _ := LoadTodos()
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			SaveTodos(todos)
			fmt.Println("✅ Todo marked as completed!")
			return
		}
	}
	fmt.Println("❌ Todo not found!")
}

// Delete a todo
func DeleteTodo(id int) {
	todos, _ := LoadTodos()
	var newTodo []Todo
	for _, todo := range todos {
		if todo.ID != id {
			newTodo = append(newTodo, todo)
		}
	}
	SaveTodos(newTodo)
	fmt.Println("✅ Todo deleted successfully!")
}
