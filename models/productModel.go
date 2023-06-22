package models

type Product struct {
	IdProduct   uint    `json:"IdProduct" gorm:"column:id_product;primaryKey;autoIncrement;"`
	NameProduct string  `json:"NameProduct" gorm:"column:name_product;"`
	Price       float64 `json:"Price" gorm:"column:price;"`
}

func (Product) TableName() string {
	return "product"
}
