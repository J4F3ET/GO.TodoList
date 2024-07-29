package main

import (
	"log"
	"net/http"

	"todo-app/internal/usecase"
	"todo-app/pkg/adapter/db"
	"todo-app/pkg/adapter/handler"

	"github.com/gorilla/mux"
)

func main(){
	repo,err := db.NewTaskPostgresRepository()
	if err != nil {
		panic(err)
	}
	taskUseCase := usecase.NewTaskUseCase(repo)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	router := mux.NewRouter()
	router.HandleFunc("/api/task/{limit:[0-9]+}/{page:[0-9]+}", taskHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/task/{id:[0-9]+}", taskHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/task", taskHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/task/{id:[0-9]+}", taskHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/task/{id:[0-9]+}", taskHandler.Delete).Methods(http.MethodDelete)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}