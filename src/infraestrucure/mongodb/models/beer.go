package mdbmodels

type Beer struct {
	ID       string  `bson:"id,omitempty" `
	DomainID string  `bson:"domain_id,omitempty"`
	Name     string  `bson:"name,omitempty"`
	Brewery  string  `bson:"brewery,omitempty"`
	Country  string  `bson:"country,omitempty"`
	Price    float32 `bson:"price,omitempty"`
	Currency string  `bson:"currency,omitempty"`
}
