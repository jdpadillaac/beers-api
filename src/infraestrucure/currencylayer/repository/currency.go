package crlayerrepo

import "C"
import (
	"encoding/json"
	"github.com/jdpadillaac/beers-api/src/domain/entity"
	"github.com/jdpadillaac/beers-api/src/domain/repository"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/currencylayer"
	crlayermodels "github.com/jdpadillaac/beers-api/src/infraestrucure/currencylayer/models"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type clCurrencyRepo struct {
	config *crlayer.Config
}

func NewClCurrencyRepo() repository.Currency {
	return &clCurrencyRepo{
		config: crlayer.NewConfigFromEnv(),
	}
}

func (c clCurrencyRepo) GetSupported() ([]entity.CurrencyAvailable, error) {
	var responseModel crlayermodels.CurrencyList
	resultList := make([]entity.CurrencyAvailable, 0)

	url := c.config.Url + "list?access_key=" + c.config.ApiKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resultList, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resultList, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resultList, err
	}

	err = json.Unmarshal(body, &responseModel)
	if err != nil {
		return resultList, err
	}

	for key, element := range responseModel.Currencies {
		newItem := entity.CurrencyAvailable{
			Name: element,
			Code: key,
		}
		resultList = append(resultList, newItem)
	}

	return resultList, nil
}

func (c clCurrencyRepo) ValidateCurrency(code string) (entity.CurrencyAvailable, error) {
	currencyList, err := c.GetSupported()
	if err != nil {
		return entity.CurrencyAvailable{}, err
	}

	for _, v := range currencyList {
		if v.Code == code {
			return v, nil
		}
	}

	return entity.CurrencyAvailable{}, errors.New("La moneda ingresada no es válida")
}
