package main

import (
	"go_beers/src/domain/models"
	"go_beers/src/infraestrucure/rest"
	"log"
	"os"
)

func main()  {
	setLogFile()
	appConfig := models.AppConfig{AppPort: "6940"}
	rest.Init(&appConfig)
}

func setLogFile() {
	file, err := os.OpenFile("assets/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}