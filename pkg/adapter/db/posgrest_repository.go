package db

import (
	"todo-app/internal/domain/entity"
	"todo-app/internal/repository"
)

type TaskPostgresRepository struct {
	conn *SingletonConnPostgres
}
func NewTaskPostgresRepository() (repository.TaskRepository,error) {
	conn,err := getInstance()
	if err != nil{
		return nil,err
	}
	return &TaskPostgresRepository{conn: conn},nil
}
// Create implements repository.TaskRepository.
func (t *TaskPostgresRepository) Create(task *entity.Task) (*entity.Task, error) {

	panic("unimplemented")
}

// Delete implements repository.TaskRepository.
func (t *TaskPostgresRepository) Delete(id int64) error {
	panic("unimplemented")
}

// GetAll implements repository.TaskRepository.
func (t *TaskPostgresRepository) GetAll() ([]*entity.Task, error) {
	panic("unimplemented")
}

// GetById implements repository.TaskRepository.
func (t *TaskPostgresRepository) GetById(id int64) (*entity.Task, error) {
	panic("unimplemented")
}

// Update implements repository.TaskRepository.
func (t *TaskPostgresRepository) Update(task *entity.Task) (*entity.Task, error) {
	panic("unimplemented")
}


