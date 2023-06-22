package models

import "gorm.io/gorm"

type ReviewProduct struct {
	IdReview   uint   `json:"IdReview" gorm:"column:id_review;primaryKey;autoIncrement;"`
	IdProduct  uint   `json:"NameProduct" gorm:"column:id_product;"`
	IdMember   uint   `json:"IdMember" gorm:"column:id_member;"`
	DescReview string `json:"DescReview" gorm:"column:desc_review;"`
}

func (ReviewProduct) TableName() string {
	return "review_product"
}

type LikeReview struct {
	IdReview uint `json:"IdReview" gorm:"column:id_review;"`
	IdMember uint `json:"IdMember" gorm:"column:id_member;"`
}

func (LikeReview) TableName() string {
	return "like_review"
}

type Review struct {
	Username         string
	Gender           string
	SkinType         string
	SkinColor        string
	DescReview       string
	JumlahLikeReview int `gorm:"column:jumlah_like_review;"`
}

type Reviews []Review

func (r *Reviews) GetReview(db *gorm.DB, Id int) error {
	return db.Table("review_product").
		Select("member.username, member.gender, member.skin_type, member.skin_color, review_product.desc_review, COUNT(like_review.id_review) as jumlah_like_review").
		Joins("JOIN member ON member.id_member = review_product.id_member").
		Joins("LEFT JOIN like_review ON like_review.id_review = review_product.id_review").
		Where("review_product.id_product = ?", Id).
		Group("review_product.id_product, member.username").
		Scan(&r).Error
}

func (l *LikeReview) Like(db *gorm.DB) error {
	return db.Model(LikeReview{}).Create(l).Error
}

func (l *LikeReview) CancelLike(db *gorm.DB) error {
	return db.Model(LikeReview{}).Where("id_review = ? ", l.IdReview).Where("id_member = ? ", l.IdMember).Delete(l).Error
}
