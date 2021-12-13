package repository

import "github.com/jdpadillaac/beers-api/src/domain/entity"

type Currency interface {
	GetSupported() ([]entity.CurrencyAvailable, error)
	ValidateCurrency(code string) (entity.CurrencyAvailable, error)
	GetUsdValue(code string) (entity.CurrencyUsdValue, error)
}
