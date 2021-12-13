package dto

type RqBeerRegister struct {
	Name     string  `json:"name" validate:"required"`
	Brewery  string  `json:"brewery" validate:"required"`
	Country  string  `json:"country" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
}
