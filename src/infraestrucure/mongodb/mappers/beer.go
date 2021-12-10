package mdbmappers

import (
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb/models"
)

type Beer struct{}

func NewBeer() *Beer {
	return &Beer{}
}

func (b Beer) FromEntity(e entity.Beer) mdbmodels.Beer {
	return mdbmodels.Beer{
		DomainID: e.ID,
		Name:     e.Name,
		Brewery:  e.Brewery,
		Country:  e.Country,
		Price:    e.Price,
		Currency: e.Currency,
	}
}
