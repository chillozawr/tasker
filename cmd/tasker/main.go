package main

import (
	"fmt"
	"time"

	"github.com/chillozawr/tasker/internal/task"
)

func main() {
	fmt.Println("Tasker started!")
	t := task.Task{ID: 1, Title: "Задача 1", Done: false, CreatedAt: time.Now()}

	if err := t.Rename("Task 1"); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(t)
}