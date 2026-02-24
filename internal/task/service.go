package task

import (
	"strings"
	"time"
)

type Service struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(title string) (Task, error) {
	title = strings.TrimSpace(title)

	if len(title) < 4 {
		return Task{}, ErrTaskTitleTooShort
	}

	t := Task{CreatedAt: time.Now()}

	if err := t.Rename(title); err != nil {
		return Task{}, err
	}

	return s.repo.Create(t)
}

func (s *Service) List() ([]Task, error) {
	return s.repo.List()
}

func (s *Service) Complete(id int) error {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	t.Complete()

	return s.repo.Update(t)
}

func (s *Service) Rename(id int, title string) error {
	title = strings.TrimSpace(title)

	if len(title) < 4 {
		return ErrTaskTitleTooShort
	}

	t, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	err = t.Rename(title)

	if err != nil {
		return err
	}

	return s.repo.Update(t)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
