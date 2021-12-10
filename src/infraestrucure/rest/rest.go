package rest

import (
	"github.com/jdpadillaac/beers-api/src/domain/models"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/rest/handlers"
	"github.com/labstack/echo/v4"
	"log"
)

func Init(c *models.AppConfig) {
	e := echo.New()

	handlers.ListenAnServe(e)
	log.Println(c.Name+": aplicacion lista y ejecuntandose en puerto", c.Port)
	e.Logger.Fatal(e.Start(":" + c.Port))
}
