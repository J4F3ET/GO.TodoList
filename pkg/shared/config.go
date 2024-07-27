package shared

import (
	"log"
	"os"
	"errors"
	"github.com/joho/godotenv"
)
var ErrNotFound = errors.New("environment variables not found")
func Init(){
	if err:=godotenv.Load();err != nil{
		log.Fatalln(err)
	}
}
//EnvGet return environment variables
func EnvGet(nameEnv string)(string,error){
	if env,ok := os.LookupEnv(nameEnv); ok {
		return env,nil
	}
	return "",ErrNotFound
}