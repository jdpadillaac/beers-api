package handlers

import (
	"github.com/jdpadillaac/beers-api/src/domain/models"
	"github.com/labstack/echo/v4"
)

type handlersConfig struct {
	echo      *echo.Echo
	baseUrl   string
	appConfig *models.AppConfig
}

func ListenAnServe(e *echo.Echo, config *models.AppConfig) {
	handlers := handlersConfig{
		echo:      e,
		baseUrl:   config.BaseUrl,
		appConfig: config,
	}

	handlers.currency()
	handlers.beer()
}
