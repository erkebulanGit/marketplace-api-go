package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"not null"`
	UserID    uint   `json:"user_id" gorm:"not null"`
	Rating    int    `json:"rating" gorm:"not null"`
	Comment   string `json:"comment"`
	User      User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
