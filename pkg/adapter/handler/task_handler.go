package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/internal/usecase"

	"github.com/gorilla/mux"
)

type TaskHandler struct{
	rep  *usecase.TaskUseCase
}
func NewTaskHandler(u *usecase.TaskUseCase)*TaskHandler{
	return &TaskHandler{rep:u}
}
func(r *TaskHandler) GetAll(res http.ResponseWriter,req *http.Request){
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
}
func(r *TaskHandler) GetById(res http.ResponseWriter,req *http.Request){
	parameters := mux.Vars(req)

	id,err := strconv.ParseInt(parameters["id"],10,64)
	if err != nil{
		http.Error(res,"invalid format id",http.StatusBadRequest)
	}

	task,err := r.rep.GetTaskByID(id)
	if err != nil{
		http.Error(res,err.Error(),http.StatusNotFound)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(task)
}
func(r *TaskHandler) Delete(res http.ResponseWriter,req *http.Request){
	var input struct{
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err!= nil {
		http.Error(res,"invalid format id",http.StatusBadRequest)
	}

	err := r.rep.DeleteTask(input.id)
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

	id,err := strconv.ParseInt(parameters["id"],10,64)
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
}