package product

import (
	"github.com/MaestroShifu/golang-fiber/src/domain/contracts"
	"github.com/MaestroShifu/golang-fiber/src/domain/entities"
)

type ProductCreate struct {
	productContract contracts.ProductContract
}

func NewProductCreate(productCont contracts.ProductContract) *ProductCreate {
	return &ProductCreate{
		productContract: productCont,
	}
}

func (pc *ProductCreate) Execute(args contracts.CreateOrUpdateProduct) (*entities.Product, error) {
	return pc.productContract.Create(args)
}
