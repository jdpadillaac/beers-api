package main

import (
	"github.com/jdpadillaac/beers-api/src/domain/models"
	mdb "github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/rest"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	setLogFile()
	envLoad()
	appConfig := models.AppConfig{
		Port:           "6940",
		MongoCnnString: os.Getenv("MONGO_CNN"),
		Name:           "BEERAPP",
		MongoDbName:    "beers-bd",
		BaseUrl:        os.Getenv("BASE_URL"),
	}
	mdb.Init(&appConfig)
	rest.Init(&appConfig)
}

func setLogFile() {
	file, err := os.OpenFile("assets/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
