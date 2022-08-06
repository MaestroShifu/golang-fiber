package adapters

import (
	"github.com/MaestroShifu/golang-fiber/src/domain/contracts"
	"github.com/MaestroShifu/golang-fiber/src/domain/entities"
	"github.com/MaestroShifu/golang-fiber/src/infrastructure/database/models"
	"gorm.io/gorm"
)

type PgsqlProduct struct {
	contracts.ProductContract
	db *gorm.DB
}

func NewPgsqlProduct(db *gorm.DB) *PgsqlProduct {
	return &PgsqlProduct{
		db: db,
	}
}

func (pgsqlProduct *PgsqlProduct) Create(args contracts.CreateOrUpdateProduct) (*entities.Product, error) {
	product := models.ProductModel{
		Sku:         args.Sku,
		Name:        args.Name,
		Description: args.Description,
		State:       string(args.State),
		Urls:        args.Urls,
		SellPrice:   args.SellPrice,
		Cost:        args.Cost,
		Barcode:     args.Barcode,
	}
	result := pgsqlProduct.db.Create(&product)
	return &entities.Product{
		Id:          product.UUID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		State:       entities.State(product.State),
		Urls:        product.Urls,
		SellPrice:   product.SellPrice,
		Cost:        product.Cost,
		Barcode:     product.Barcode,
	}, result.Error
}
