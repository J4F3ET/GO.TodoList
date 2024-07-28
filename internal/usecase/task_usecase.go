package usecase

import (
	"errors"
	"todo-app/internal/domain/entity"
	"todo-app/internal/repository"
)

var ErrNotFound = errors.New("task not found")

// TaskUseCase handles the business logic related to the task
type TaskUseCase struct {
	repo repository.TaskRepository
}
// NewTaskUseCase Return new instance to TaskUseCase
func NewTaskUseCase(repository repository.TaskRepository) *TaskUseCase{
	return &TaskUseCase{repo: repository}
}
// GetAll Return all task or error
func (uc *TaskUseCase) GetTasks(limit, page int)([]*entity.Task, error){
	if page < 0{
		page = page*(-1)
	}
	if page == 0{
		page = 1
	}
	return uc.repo.GetAll(limit, page-1)
}
// GetTaskByID Return Task corresponding to the ID or error
func (uc *TaskUseCase) GetTaskByID(id int64)(*entity.Task,error){
	task, err := uc.repo.GetById(id)
	if err != nil{
		return nil,ErrNotFound
	}
	return task,nil
}
// CreateTask Created new task and return the new task
func (uc *TaskUseCase) CreateTask(title string)(*entity.Task,error){
	task := &entity.Task{
		Title: title,
		Completed: false,
	}
	return uc.repo.Create(task)
}
// UpdateTask Updates all parameters of a task
func (uc *TaskUseCase) UpdateTask(id int64,title string,completed bool)(*entity.Task,error){
	task,err := uc.repo.GetById(id)
	if err != nil{
		return nil,ErrNotFound
	}
	task.Title = title
	task.Completed = completed
	return uc.repo.Update(task)
}
// CompleteTask Update only parameter 'completed' with value 'true' by default
func (uc *TaskUseCase) CompleteTask(id int64)(*entity.Task,error){
	task,err := uc.repo.GetById(id)
	if err != nil{
		return nil,ErrNotFound
	}
	task.Completed = true
	return uc.repo.Update(task)
}
//DeleteTask Delete task
func (uc *TaskUseCase) DeleteTask(id int64)(error){
	return uc.repo.Delete(id)
}
