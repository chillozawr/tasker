package task

import (
	"errors"
	"time"
)

var ErrTaskTitleTooShort = errors.New("Заголовок слишком короткий")

type Task struct {
	ID int
	Title string
	Done bool
	CreatedAt time.Time
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Rename(title string) error {
	if (len(title) < 4) {
		return ErrTaskTitleTooShort
	}

	t.Title = title
	return nil
}

func (t Task) IsDone() bool {
	return t.Done
}