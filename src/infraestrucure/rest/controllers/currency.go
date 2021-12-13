package controllers

import (
	"github.com/jdpadillaac/beers-api/src/app/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Currency struct {
	useCase *usecase.Currency
}

func NewCurrency(uc *usecase.Currency) *Currency {
	return &Currency{
		useCase: uc,
	}
}

func (cr Currency) GetSupported(c echo.Context) error {

	result, err := cr.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, newResponse(true, requestSuccessMsg, result))
}
