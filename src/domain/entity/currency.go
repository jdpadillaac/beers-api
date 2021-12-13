package entity

type CurrencyAvailable struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CurrencyUsdValue struct {
	Code  string  `json:"code"`
	Value float32 `json:"value"`
}
