package task

type MemoryRepository struct {
	tasks map[int]Task
	nextID int
}



func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks: make(map[int]Task),
		nextID: 1,
	}
}

func (mr *MemoryRepository) Create(task Task) (Task, error) {
	task.ID = mr.nextID
	mr.tasks[mr.nextID] = task

	mr.nextID++

	return task, nil
}

func (mr *MemoryRepository) List() ([]Task, error) {
	list := make([]Task, 0, len(mr.tasks))

	for _, task := range mr.tasks {
		list = append(list, task)
	}

	return list, nil
}

func (mr *MemoryRepository) GetByID(id int) (Task, error) {
	if task, ok := mr.tasks[id]; ok {
		return task, nil
	}

	return Task{}, ErrTaskNotFound
}

func (mr *MemoryRepository) Update(task Task) error {
	if _, ok := mr.tasks[task.ID]; ok {
		mr.tasks[task.ID] = task
		return nil
	}
	
	return ErrTaskNotFound
}

func (mr *MemoryRepository) Delete(id int) error {
	if _, ok := mr.tasks[id]; ok {
		delete(mr.tasks, id)
		return nil
	}
	
	return ErrTaskNotFound
}