package controllers

import (
	"net/http"

	"github.com/MaestroShifu/golang-fiber/src/application/product"
	"github.com/MaestroShifu/golang-fiber/src/domain/contracts"
	"github.com/MaestroShifu/golang-fiber/src/domain/entities"
	"github.com/MaestroShifu/golang-fiber/src/infrastructure/interface/adapters"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateOrUpdateProduct struct {
	Sku         string   `json:"sku" xml:"sku" form:"sku"`
	Name        string   `json:"name" xml:"name" form:"name"`
	Description string   `json:"description" xml:"description" form:"description"`
	State       string   `json:"state" xml:"state" form:"state"`
	Urls        []string `json:"urls" xml:"urls" form:"urls"`
	SellPrice   float32  `json:"sellPrice" xml:"sellPrice" form:"sellPrice"`
	Cost        float32  `json:"cost" xml:"cost" form:"cost"`
	Barcode     string   `json:"barcode" xml:"barcode" form:"barcode"`
}

func AddProductController(app *fiber.App, db *gorm.DB) {

	app.Post("/products", func(c *fiber.Ctx) error {
		args := new(CreateOrUpdateProduct)
		err := c.BodyParser(args)

		if err != nil {
			return err
		}

		pgsqlProduct := adapters.NewPgsqlProduct(db)
		useCase := product.NewProductCreate(pgsqlProduct)
		product, _ := useCase.Execute(contracts.CreateOrUpdateProduct{
			Sku:         args.Sku,
			Name:        args.Name,
			Description: args.Description,
			State:       entities.State(args.State),
			Urls:        args.Urls,
			SellPrice:   args.SellPrice,
			Cost:        args.Cost,
			Barcode:     args.Barcode,
		})

		return c.Status(http.StatusOK).JSON(product)
	})
}
