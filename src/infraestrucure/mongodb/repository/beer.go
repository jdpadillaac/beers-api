package mdbrepo

import (
	"context"
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/domain/repository"
	mdb "github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb/mappers"
	"go.mongodb.org/mongo-driver/mongo"
)

type beerRepo struct {
	collection *mongo.Collection
	mapper     *mdbmappers.Beer
}

func newBeerRepo() repository.BeerRepository {
	return &beerRepo{
		collection: mdb.GetCollection("beers"),
		mapper:     mdbmappers.NewBeer(),
	}
}

func (b beerRepo) Save(beer entity.Beer) error {
	model := b.mapper.FromEntity(beer)

	_, err := b.collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}
