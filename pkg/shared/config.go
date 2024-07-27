package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func Init(){
	if err:=godotenv.Load();err != nil{
		log.Fatalln(err)
	}
}
func EnvGet(nameEnv,nameDefault string)string{
	env,ok := os.LookupEnv(nameEnv)
	if !ok {
		return nameDefault
	}
	return env
}