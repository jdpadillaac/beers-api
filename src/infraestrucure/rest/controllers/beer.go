package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jdpadillaac/beers-api/src/app/dto"
	"github.com/jdpadillaac/beers-api/src/app/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type beerCtr struct {
	val     *validator.Validate
	useCase *usecase.Beer
}

func NewBeerCtr(useCase *usecase.Beer) *beerCtr {
	return &beerCtr{val: validator.New(), useCase: useCase}
}

func (b beerCtr) GetAll(c echo.Context) error {
	result, err := b.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, newResponse(false, requestSuccessMsg, result))
}

func (b beerCtr) GetBoxPrice(c echo.Context) error {
	finalQuantity := 0

	currency := c.QueryParam("currency")
	quantity := c.QueryParam("quantity")
	ID := c.Param("beerID")

	if len(quantity) > 0 {
		intValue, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, newResponse(
				false,
				"El número no es un entero válido "+err.Error(),
				nil,
			))
		}
		finalQuantity = int(intValue)
	}
	model := dto.RqCalculateBoxPrice{
		BeerID:       ID,
		Quantity:     finalQuantity,
		CurrencyCode: currency,
	}

	result, err := b.useCase.BoxPrice(model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, newResponse(true, requestSuccessMsg, result))
}

func (b beerCtr) GetByID(c echo.Context) error {
	ID := c.Param("ID")

	result, err := b.useCase.FindByID(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, newResponse(true, requestSuccessMsg, result))
}

func (b beerCtr) Post(c echo.Context) error {
	var model dto.RqBeerRegister

	err := c.Bind(&model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(false, err.Error(), nil))
	}

	err = b.val.Struct(model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse(false, err.Error(), nil))
	}

	err = b.useCase.Save(model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newResponse(false, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, newResponse(true, registerSuccessMsg, nil))
}
