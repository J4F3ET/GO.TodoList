package db

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"todo-app/pkg/shared"

	_ "github.com/lib/pq"
)
var ErrGetInstance = errors.New("error instance conn to database")
// Lock variable to access goroutine
var lock = &sync.Mutex{}

type SingletonConnPostgres struct{
	db *sql.DB
}
var connPostgres *SingletonConnPostgres

func getInstance() (*SingletonConnPostgres,error){
	if connPostgres == nil {
		lock.Lock()
		defer lock.Unlock()

		connStr,err := getConnString()
		if err != nil{
			return nil,ErrGetInstance
		}
		connPostgres,err = createInstance(connStr)
		if err != nil{
			return nil,ErrGetInstance
		}
	}
	return connPostgres,nil
}
func createInstance(connStr string)(*SingletonConnPostgres,error){
	conn,err := sql.Open("postgres",connStr)
	if err != nil{
		return nil,err
	}
	conn.SetMaxOpenConns(3)//Conn max open(active)
	conn.SetMaxIdleConns(7)// Conn min idle(inactive)
	//conn.SetMaxIdleConns()
	//conn.SetMaxOpenConns()
	return &SingletonConnPostgres{db:conn},nil
}
func getConnString()(string,error){
	envVars := map[string]string{
                "DB_USERNAME": "",
                "DB_PASSWORD": "",
                "DB_HOST":     "",
                "DB_PORT":     "",
                "DB_NAME":     "",
        }

	for key := range envVars {
		value, err := shared.EnvGet(key)
		if err != nil {
			return "", err
		}
		envVars[key] = value
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envVars["DB_USERNAME"],
		envVars["DB_PASSWORD"],
		envVars["DB_HOST"],
		envVars["DB_PORT"],
		envVars["DB_NAME"],
	)
	return connStr,nil
}