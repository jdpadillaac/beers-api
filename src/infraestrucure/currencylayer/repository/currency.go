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
	"strings"
)

type clCurrencyRepo struct {
	config *crlayer.Config
}

func NewClCurrencyRepo() repository.Currency {
	return &clCurrencyRepo{
		config: crlayer.NewConfigFromEnv(),
	}
}

func (c clCurrencyRepo) GetUsdValue(code string) (entity.CurrencyUsdValue, error) {
	var respModel crlayermodels.CurrencyValue
	var model entity.CurrencyUsdValue

	url := c.config.Url + "live?access_key=" + c.config.ApiKey + "&currencies=" + code

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model, err
	}

	err = json.Unmarshal(body, &respModel)
	if err != nil {
		return model, err
	}

	for key, element := range respModel.Quotes {
		codeSplit := strings.Split(key, "USD")
		model = entity.CurrencyUsdValue{
			Code:  codeSplit[1],
			Value: element,
		}
	}

	return model, nil
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
