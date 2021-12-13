package dto

type RqBeerRegister struct {
	Name     string  `json:"name" validate:"required"`
	Brewery  string  `json:"brewery" validate:"required"`
	Country  string  `json:"country" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
}

type RqCalculateBoxPrice struct {
	BeerID       string `json:"beer_id"`
	Quantity     int    `json:"quantity"`
	CurrencyCode string `json:"currency_code"`
}

type RsCalculatedBoxPrice struct {
	TotalPrice float32 `json:"total_price"`
}
