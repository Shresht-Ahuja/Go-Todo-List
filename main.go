package main

import "fmt"

const file = "todos.json"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	fmt.Println("Your Todo List:")
	todos.print()
	storage.Save(todos)
}
