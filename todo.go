package main

import ("fmt"
		"time"
		"os"
		"strconv"
		"github.com/aquasecurity/table"
)

type Todo struct{
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
	Title: title,
	Completed: false,
	CompletedAt: nil,
	CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) list() {
    for i, todo := range *todos {
        status := "Pending"
        if todo.Completed {
            status = "Completed"
        }
        fmt.Printf("%d. %s [%s]\n", i+1, todo.Title, status)
    }
}

func (todos *Todos) delete(index int) {
	// checking the index exists or not
    if index < 0 || index >= len(*todos) {
        fmt.Println("Task at index %d does not exist", index)
        return
    }
    *todos = append((*todos)[:index], (*todos)[index+1:]...)
}

func (todos *Todos) toggle(index int) {
	// checking the index exists or not
	if index < 0 || index >= len(*todos) {
		fmt.Println("Task does not exist")
		return
	}

	todo := &(*todos)[index]

	if todo.Completed {
		todo.Completed = false
		todo.CompletedAt = nil
	} else {
		todo.Completed = true
		now := time.Now()
		todo.CompletedAt = &now
	}

	fmt.Printf("Task '%s' is completed yay!!!\n", (*todos)[index].Title)
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Created At", "Completed At")
	for index, todo := range *todos {
        status := "❌"
		completedAt := ""
        if todo.Completed {
            status = "✅"
			completedAt = todo.CompletedAt.Format(time.RFC1123)
        }
	table.AddRow(strconv.Itoa(index), todo.Title, status, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}