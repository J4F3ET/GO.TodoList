package repository

import "todo-app/internal/entity"

type TaskRepository interface{
	GetAll(limit, offset int)(map[uint64]*entity.Task,error)
	GetById(id uint64)(*entity.Task,error)
	Create(task *entity.Task)(*entity.Task,error)
	Update(task *entity.Task)(*entity.Task,error)
	Delete(id uint64)(error)
}