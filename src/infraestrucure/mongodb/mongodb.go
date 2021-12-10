package mdb

import (
	"context"
	"github.com/jdpadillaac/beers-api/src/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var db *mongo.Database

func Init(app *models.AppConfig) {
	opts := options.Client().ApplyURI(app.MongoCnnString)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(app.Name+": Error en conexión a la base de datos mongodb", err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(app.Name+": error en ping a la base de datos", err.Error())
	}

	db = client.Database(app.MongoDbName)
	log.Println(app.Name + ": Base de datos mongo db conectada y lista para usarse")
}

func GetCollection(name string) *mongo.Collection {
	if db == nil {
		log.Fatal("No se estableció conexión con la base de datos")
	}
	return db.Collection(name)
}
