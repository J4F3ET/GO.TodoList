package repository

import "todo-app/internal/domain/entity"

type TaskRepository interface{
	GetAll(limit, offset int)([]*entity.Task,error)
	GetById(id int64)(*entity.Task,error)
	Create(task *entity.Task)(*entity.Task,error)
	Update(task *entity.Task)(*entity.Task,error)
	Delete(id int64)(error)
}