package models

import (
	"time"

	"github.com/lib/pq"
)

type ProductModel struct {
	UUID        string         `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()"`
	Sku         string         `gorm:"column:sku;not null"`
	Name        string         `gorm:"column:name;not null"`
	Description string         `gorm:"column:description;not null"`
	State       string         `gorm:"column:state"`
	Urls        pq.StringArray `gorm:"column:urls;type:text[];default:array[]::text[]"`
	SellPrice   float32        `gorm:"column:sell_price;not null"`
	Cost        float32        `gorm:"column:cost"`
	Barcode     string         `gorm:"column:bar_code"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime:true"`
	UpdateAt    time.Time      `gorm:"column:update_at;autoUpdateTime:true"`
}

func (ProductModel) TableName() string {
	return "products"
}
