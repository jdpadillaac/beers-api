package usecase

import (
	"github.com/google/uuid"
	"github.com/jdpadillaac/beers-api/src/app/dto"
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/domain/repository"
	"github.com/pkg/errors"
	"strings"
)

type Beer struct {
	repository   repository.BeerRepository
	currencyRepo repository.Currency
}

func NewBeer(repository repository.BeerRepository, currencyRepo repository.Currency) *Beer {
	return &Beer{repository: repository, currencyRepo: currencyRepo}
}

func (b Beer) GetAll() ([]entity.Beer, error) {
	result, err := b.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	return result, nil
}

func (b Beer) BoxPrice(model dto.RqCalculateBoxPrice) (dto.RsCalculatedBoxPrice, error) {
	var respModel dto.RsCalculatedBoxPrice
	var beerQuantity int
	var beerCurrency string

	if model.Quantity == 0 {
		beerQuantity = 6
	} else {
		beerQuantity = model.Quantity
	}

	if len(model.CurrencyCode) == 0 {
		beerCurrency = "USD"
	} else {
		beerCurrency = model.CurrencyCode
	}

	beer, err := b.repository.FindByID(model.BeerID)
	if err != nil {
		return respModel, errors.Wrap(err, "Error en validación en la base de datos")
	}

	requestCurrency, err := b.currencyRepo.GetUsdValue(strings.ToUpper(beerCurrency))
	if err != nil {
		return respModel, errors.Wrap(err, "Error en validación de moneda")
	}

	recordCurrency, err := b.currencyRepo.GetUsdValue(strings.ToUpper(beer.Currency))
	if err != nil {
		return respModel, errors.Wrap(err, "Error en validación de moneda")
	}

	totalQuantity := beer.Price * float32(beerQuantity)
	requiredQuantity := totalQuantity / recordCurrency.Value
	requestValue := requiredQuantity * requestCurrency.Value

	respModel = dto.RsCalculatedBoxPrice{TotalPrice: requestValue}

	return respModel, nil
}

func (b Beer) Save(model dto.RqBeerRegister) error {
	crr, err := b.currencyRepo.ValidateCurrency(strings.ToUpper(model.Currency))
	if err != nil {
		return errors.Wrap(err, "Error en validación de moneda")
	}

	beer := entity.Beer{
		ID:       uuid.New().String(),
		Name:     model.Name,
		Brewery:  model.Brewery,
		Country:  model.Country,
		Price:    model.Price,
		Currency: crr.Code,
	}

	err = b.repository.Save(beer)
	if err != nil {
		return errors.Wrap(err, "Error al guardar en la base de datos")
	}

	return nil
}

func (b Beer) FindByID(ID string) (entity.Beer, error) {
	result, err := b.repository.FindByID(ID)
	if err != nil {
		return result, errors.Wrap(err, "Error en validación de documento en la base de datos")
	}

	return result, nil
}
