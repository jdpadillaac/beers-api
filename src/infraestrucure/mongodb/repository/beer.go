package mdbrepo

import (
	"context"
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/domain/repository"
	mdb "github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb/mappers"
	mdbmodels "github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type beerRepo struct {
	collection *mongo.Collection
	mapper     *mdbmappers.Beer
}

func NewBeerRepo() repository.BeerRepository {
	return &beerRepo{
		collection: mdb.GetCollection("beers"),
		mapper:     mdbmappers.NewBeer(),
	}
}

func (b beerRepo) GetAll() ([]entity.Beer, error) {
	resultList := make([]entity.Beer, 0)

	cursor, err := b.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var itemModel mdbmodels.Beer

		decodeErr := cursor.Decode(&itemModel)
		if decodeErr != nil {
			return resultList, err
		}

		newItem := b.mapper.FromModel(itemModel)
		resultList = append(resultList, newItem)
	}

	return resultList, nil
}

func (b beerRepo) FindByID(ID string) (entity.Beer, error) {
	var model mdbmodels.Beer
	var resultModel entity.Beer

	filter := bson.M{"domain_id": ID}
	err := b.collection.FindOne(context.Background(), filter).Decode(&model)
	if err != nil {
		return resultModel, err
	}

	resultModel = b.mapper.FromModel(model)
	return resultModel, nil
}

func (b beerRepo) Save(beer entity.Beer) error {
	model := b.mapper.FromEntity(beer)

	_, err := b.collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}
