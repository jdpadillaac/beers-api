package rest

import (
	"github.com/labstack/echo/v4"
	"go_beers/src/domain/models"
	"go_beers/src/infraestrucure/rest/handlers"
	"log"
)

func Init(c *models.AppConfig)  {
	e := echo.New()

	handlers.ListenAnServe(e)
	log.Println("BEERSAPI: aplicacion lista y ejecuntandose en puerto", c.AppPort)
	e.Logger.Fatal(e.Start(":" + c.AppPort))
}