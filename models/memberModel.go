package models

import "gorm.io/gorm"

type Member struct {
	IdMember  uint   `json:"IdMember" gorm:"column:id_member;primaryKey;autoIncrement;"`
	Username  string `json:"Username" gorm:"column:username;"`
	Gender    string `json:"Gender" gorm:"column:gender;"`
	SkinType  string `json:"SkinType" gorm:"column:skin_type;"`
	SkinColor string `json:"SkinColor" gorm:"column:skin_color;"`
}

func (Member) TableName() string {
	return "member"
}

type Members []Member

func (m *Member) Create(db *gorm.DB) error {
	return db.Model(Member{}).Create(m).Error
}

func (m *Member) GetById(db *gorm.DB, Id int) error {
	return db.Model(Member{}).Where("id_member = ?", Id).First(m).Error
}

func (m *Member) Update(db *gorm.DB) error {
	return db.Model(Member{}).Omit("id_member").Where("id_member = ?", m.IdMember).Updates(m).Error
}

func (m *Member) Delete(db *gorm.DB) error {
	return db.Model(Member{}).Where("id_member = ?", m.IdMember).Delete(m).Error
}

func (m *Members) GetAll(db *gorm.DB) error {
	return db.Model(Member{}).Find(m).Error
}
