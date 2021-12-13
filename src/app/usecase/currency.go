package usecase

import (
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/domain/repository"
	"github.com/pkg/errors"
)

type Currency struct {
	repo repository.Currency
}

func NewCurrency(repo repository.Currency) *Currency {
	return &Currency{repo: repo}
}

func (c Currency) GetAll() ([]entity.CurrencyAvailable, error) {
	result, err := c.repo.GetSupported()

	if err != nil {
		return nil, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	return result, nil
}
