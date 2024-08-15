package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"log"
	"todo-app/internal/entity"
	"todo-app/internal/usecase"
	"github.com/gorilla/mux"
)

type TaskHandler struct{
	rep  *usecase.TaskUseCase
}
func NewTaskHandler(u *usecase.TaskUseCase)*TaskHandler{
	return &TaskHandler{rep:u}
}
// GetAll retrieves all tasks with pagination
// @Summary Get all tasks
// @Description Retrieve all tasks with pagination
// Tags Task
// @Accept  json
// @Produce  json
// @Param limit path int true "Limit"
// @Param page path int true "Page"
// @Success 200 {objetc} map[uint64]Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/{limit}/{page} [get]
func(r *TaskHandler) GetAll(res http.ResponseWriter,req *http.Request){
	start := time.Now()
	parameters := mux.Vars(req)

	limit,err := strconv.Atoi(parameters["limit"])
	if err != nil{
		http.Error(res,"invalid limit",http.StatusBadRequest)
		return
	}

	page,err := strconv.Atoi(parameters["page"])
	if err != nil{
		http.Error(res,"invalid page",http.StatusBadRequest)
		return
	}

	tasks,err := r.rep.GetTasks( limit,page)
	if err != nil{
		http.Error(res,err.Error(),http.StatusNotFound)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(tasks)
    log.Printf("\nTiempo de ejecución: %s \nNúmero de elementos: %d", time.Since(start),limit)
}
// GetById retrieves all tasks with pagination
// @Summary Get Task by id
// @Description Retrive Task
// Tags Task
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Success 200 {entity.Task} entity.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/{id} [get]
func(r *TaskHandler) GetById(res http.ResponseWriter,req *http.Request){
	parameters := mux.Vars(req)
	var err error
	id,err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil{
		http.Error(res,"invalid format id",http.StatusBadRequest)
	}
	var task *entity.Task
	task, err = r.rep.GetTaskByID(id)
	if err != nil{
		http.Error(res,err.Error(),http.StatusNotFound)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(task)
}
// Delete delete task by id
// @Summary Delete Task by id
// @Description Delete Task by id
// Tags Task
// @Accept  json
// @Param id path int true "Id"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /task/{id} [delete]
func(r *TaskHandler) Delete(res http.ResponseWriter,req *http.Request){
	var input struct{
		Id uint64 `json:"id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err!= nil {
		http.Error(res,"invalid format id",http.StatusBadRequest)
	}

	err := r.rep.DeleteTask(input.Id)
	if err != nil {
		http.Error(res,err.Error(),http.StatusNotFound)
	}

	res.WriteHeader(http.StatusNoContent)
}

func(r *TaskHandler) Update(res http.ResponseWriter,req *http.Request){
	var input struct{
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}

	parameters := mux.Vars(req)

	id,err := strconv.ParseUint(parameters["id"],10,64)
	if err != nil{
		http.Error(res,"invalid format id",http.StatusBadGateway)
	}
	if err := json.NewDecoder(req.Body).Decode(&input);err != nil{
		http.Error(res,"invalid format of the parameters",http.StatusBadRequest)
	}
	task,err := r.rep.UpdateTask(id,input.Title,input.Completed)
	if err != nil{
		http.Error(res,err.Error(),http.StatusNotFound)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(task)
}
func(r *TaskHandler) Create(res http.ResponseWriter,req *http.Request){
	start := time.Now()
	var input struct{
		Title string `json:"title"`
	}
	if err := json.NewDecoder(req.Body).Decode(&input);err != nil{
		http.Error(res,"invalid format of the parameters",http.StatusBadRequest)
	}

	task,err := r.rep.CreateTask(input.Title)
	if err != nil{
		http.Error(res,err.Error(),http.StatusNotFound)
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(task)
	elapsed := time.Since(start)
    log.Printf("Create Tiempo de ejecución: %s", elapsed)
}