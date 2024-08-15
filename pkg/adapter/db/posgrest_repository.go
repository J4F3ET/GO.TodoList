package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	"todo-app/internal/entity"
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

	const query string = `INSERT INTO task (title,completed) VALUES ($1,$2) RETURNING id`
	var id uint64

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	row := connPostgres.db.QueryRowContext(ctx,query,task.Title,task.Completed)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	task.Id = id
	return task,nil
}

// Delete implements repository.TaskRepository.
func (t *TaskPostgresRepository) Delete(id uint64) error {
	const query string = `DELETE FROM task WHERE id=$1`
	var rows int64
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	result,err := connPostgres.db.ExecContext(ctx,query,id)
	if err != nil {
		return err
	}
	rows,err = result.RowsAffected()
	if err != nil{
		return err
	}

	if rows != 0 {
		return errors.New("item with element specified ID not found")
	}else if rows > 1 {
		errMessage := fmt.Sprintf("expected single row affected, got %d rows affected", rows)
		return errors.New(errMessage)
	}

	return nil
}

var processRow = func(rows *sql.Rows, taskChan chan<- *entity.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for rows.Next() {
		task := &entity.Task{}
		err := rows.Scan(&task.Id, &task.Title, &task.Completed)
		if err != nil {
			log.Printf("Error al escanear fila: %v", err)
			continue
		}
		taskChan <- task
	}
	defer close(taskChan)
}
var receiveTask = func(taskChan chan *entity.Task ,wg *sync.WaitGroup,tasksMap map[uint64]*entity.Task){
	defer wg.Done()
	for task := range taskChan{
		tasksMap[task.Id] = task
	}
}
// GetAll implements repository.TaskRepository.
func (t *TaskPostgresRepository) GetAll(limit, offset int) (map[uint64]*entity.Task, error) {
	const query string = `
		SELECT id, title, completed
		FROM task
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	rows,err := connPostgres.db.QueryContext(ctx,query,limit,offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskChan := make(chan *entity.Task,limit/4)
	tasksMap:= make(map[uint64]*entity.Task)
	var wg sync.WaitGroup
	wg.Add(2)
	go processRow(rows,taskChan,&wg)
	go receiveTask(taskChan,&wg,tasksMap)
	wg.Wait()
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasksMap,nil
}

// GetById implements repository.TaskRepository.
func (t *TaskPostgresRepository) GetById(id uint64) (*entity.Task, error) {

	const query =  `
		SELECT title,completed
		FROM task
		WHERE id=$1
	`
	var task = &entity.Task{Id: id}

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	row := connPostgres.db.QueryRowContext(ctx,query,id)

	err := row.Scan(&task.Title, &task.Completed)
	if err != nil {
		return nil,err
	}

	return task,nil
}

// Update implements repository.TaskRepository.
func (t *TaskPostgresRepository) Update(task *entity.Task) (*entity.Task, error) {
	const query string = `UPDATE task SET title = $1, completed $2 WHERE id=$3`
	var rows int64
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	result,err := connPostgres.db.ExecContext(ctx,query,task.Title,task.Completed,task.Id)
	if err != nil {
		return nil,err
	}

	rows,err = result.RowsAffected()
	if err != nil{
		return nil,err
	}

	if rows != 0 {
		return nil,errors.New("item with element specified ID not found")
	}else if rows > 1 {
		errMessage := fmt.Sprintf("expected single row affected, got %d rows affected", rows)
		return nil,errors.New(errMessage)
	}

	return task,nil
}


