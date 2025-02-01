package main

import "fmt"

const file = "todos.json"

func main() {
	var todos Todos

	todos.add("Buy milk")
	todos.add("Build a CLI app")

	fmt.Println("Your Todo List:")
	todos.toggle(0)
	todos.print()
}
