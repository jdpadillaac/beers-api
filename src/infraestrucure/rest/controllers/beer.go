package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jdpadillaac/beers-api/src/app/dto"
	"github.com/jdpadillaac/beers-api/src/app/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type beerCtr struct {
	val     *validator.Validate
	useCase *usecase.Beer
}

func NewBeerCtr(useCase *usecase.Beer) *beerCtr {
	return &beerCtr{val: validator.New(), useCase: useCase}
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
