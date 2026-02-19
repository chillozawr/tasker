package task

import (
	"strings"
	"time"

	"github.com/chillozawr/tasker/internal/task"
)

type Service struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *Service

func Create(title string) (Task, error) {
	repo := task.NewMemoryRepository()
	title = strings.Trim(title, " ")

	if len(title) < 4 {
		return Task{}, ErrTaskTitleTooShort
	}

	return repo.Create(task.Task{Title: title, CreatedAt: time.Now()}), nil
}