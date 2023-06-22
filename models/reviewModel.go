package models

type ReviewProduct struct {
	IdReview   uint   `json:"IdProduct" gorm:"column:id_product;primaryKey;autoIncrement;"`
	IdProduct  uint   `json:"NameProduct" gorm:"column:id_product;"`
	IdMember   uint   `json:"IdMember" gorm:"column:id_member;"`
	DescReview string `json:"DescReview" gorm:"column:desc_review;"`
}

func (ReviewProduct) TableName() string {
	return "review_product"
}

type LikeReview struct {
	IdReview uint `json:"IdProduct" gorm:"column:id_product;"`
	IdMember uint `json:"IdMember" gorm:"column:id_member;"`
}

func (LikeReview) TableName() string {
	return "like_review"
}
