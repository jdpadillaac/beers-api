package repository

import "github.com/jdpadillaac/beers-api/src/domain/entity"

type BeerRepository interface {
	Save(beer entity.Beer) error
	FindByID(ID string) (entity.Beer, error)
	GetAll() ([]entity.Beer, error)
}
