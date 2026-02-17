package main

import (
	"fmt"
	"time"

	"github.com/chillozawr/tasker/internal/task"
)

func main() {
	repo := task.NewMemoryRepository()
	fmt.Println("Tasker started!")
	_, err := repo.Create(task.Task{Title: "Задача 1", CreatedAt: time.Now()})
	if err != nil {
		panic(err)
	}

	t2, err := repo.Create(task.Task{Title: "Задача 2", CreatedAt: time.Now()})
	if err != nil {
		panic(err)
	}


	tasks, err := repo.List()
	if err != nil {
		panic(err)
	}
	fmt.Println("before delete:", tasks)

	if err := repo.Delete(t2.ID); err != nil {
		fmt.Println("delete error:", err)
	}
	
	tasks, err = repo.List()
	if err != nil {
		panic(err)
	}
	fmt.Println("after delete:", tasks)
 
	
	if _, err = repo.GetByID(t2.ID); err != nil {
		fmt.Println("get deleted: ", err)
	}
}