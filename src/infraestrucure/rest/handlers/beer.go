package handlers

import (
	"github.com/jdpadillaac/beers-api/src/app/usecase"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/currencylayer/repository"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/mongodb/repository"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/rest/controllers"
)

func (c handlersConfig) beer() {
	endPoint := c.baseUrl + "beers"

	repo := mdbrepo.NewBeerRepo()
	useCase := usecase.NewBeer(repo, crlayerrepo.NewClCurrencyRepo())
	ctr := controllers.NewBeerCtr(useCase)

	g := c.echo.Group(endPoint)

	g.POST("", ctr.Post)

	g.GET("/:ID", ctr.GetByID)

	g.GET("/:beerID/boxprice", ctr.GetBoxPrice)

}
