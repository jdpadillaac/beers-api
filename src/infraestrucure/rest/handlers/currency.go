package handlers

import (
	"github.com/jdpadillaac/beers-api/src/app/usecase"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/currencylayer/repository"
	"github.com/jdpadillaac/beers-api/src/infraestrucure/rest/controllers"
)

func (c handlersConfig) currency() {
	endPoint := c.baseUrl + "currency"

	repo := crlayerrepo.NewClCurrencyRepo()
	useCase := usecase.NewCurrency(repo)
	ctr := controllers.NewCurrency(useCase)

	g := c.echo.Group(endPoint)

	g.GET("", ctr.GetSupported)

}
