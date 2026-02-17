package task

import "errors"

var ErrTaskNotFound = errors.New("Задача не найдена")

type TaskRepository interface {
	Create(task Task) (Task, error)
	List() ([]Task, error)
	GetByID(id int) (Task, error)
	Update(task Task) error
	Delete(id int) error
}