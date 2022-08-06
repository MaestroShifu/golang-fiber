package contracts

import "github.com/MaestroShifu/golang-fiber/src/domain/entities"

type CreateOrUpdateProduct struct {
	Sku         string
	Name        string
	Description string
	State       entities.State
	Urls        []string
	SellPrice   float32
	Cost        float32
	Barcode     string
}

type ProductContract interface {
	Create(args CreateOrUpdateProduct) (*entities.Product, error)
}
