package main

import (
	"log"
	"userApp/rest"
	"userApp/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.LazyEnvVariableInit()

	s := rest.NewServer()
	s.Run()
}
