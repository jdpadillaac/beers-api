package crlayermodels

type CurrencyList struct {
	Success    bool              `json:"success"`
	Terms      string            `json:"terms"`
	Privacy    string            `json:"privacy"`
	Currencies map[string]string `json:"currencies"`
}
