package models

import "gorm.io/gorm"

type Product struct {
	IdProduct   uint    `json:"IdProduct" gorm:"column:id_product;primaryKey;autoIncrement;"`
	NameProduct string  `json:"NameProduct" gorm:"column:name_product;"`
	Price       float64 `json:"Price" gorm:"column:price;"`
}

func (Product) TableName() string {
	return "product"
}

type Products []Product

func (p *Product) GetById(db *gorm.DB, Id int) error {
	return db.Model(Product{}).Where("id_product = ?", Id).First(p).Error
}
