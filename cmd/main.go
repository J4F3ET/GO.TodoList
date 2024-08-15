package main

import (
	"log"
	"net/http"

	"todo-app/internal/usecase"
	"todo-app/pkg/adapter/db"
	"todo-app/pkg/adapter/handler"

	_"todo-app/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)
// @title Todo List(go)
// @version 1.0
// @description This is project by learning go

// @contact.name Jafet
// @contact.url https://main--j4f3t.netlify.app/
// @contact.email jafetstivenlopezzuniga@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main(){
	repo,err := db.NewTaskPostgresRepository()
	if err != nil {
		panic(err)
	}
	taskUseCase := usecase.NewTaskUseCase(repo)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/task/{limit:[0-9]+}/{page:[0-9]+}", taskHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/task/{id:[0-9]+}", taskHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/task", taskHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/task/{id:[0-9]+}", taskHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/task/{id:[0-9]+}", taskHandler.Delete).Methods(http.MethodDelete)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}