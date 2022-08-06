package entities

import "errors"

type State string

const (
	ACTIVATE    State = "Activate"
	DESACTIVATE State = "Desactivate"
)

type Product struct {
	Id          string   `json:"id"`
	Sku         string   `json:"sku"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	State       State    `json:"state"`
	Urls        []string `json:"urls"`
	SellPrice   float32  `json:"sellPrice"`
	Cost        float32  `json:"cost"`
	Barcode     string   `json:"barcode"`
}

func validateState(s State) bool {
	return s == ACTIVATE || s == DESACTIVATE
}

func NewProduct(
	sku string,
	name string,
	description string,
	state State,
	urls []string,
	sellPrice float32,
	cost float32,
	barcode string,
) (*Product, error) {
	if !validateState(state) {
		return nil, errors.New("State invalid")
	}
	return &Product{
		Sku:         sku,
		Name:        name,
		Description: description,
		State:       state,
		Urls:        urls,
		SellPrice:   sellPrice,
		Cost:        cost,
		Barcode:     barcode,
	}, nil
}
