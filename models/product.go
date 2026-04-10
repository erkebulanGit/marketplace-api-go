package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description"`
	Price       float64  `json:"price" gorm:"not null"`
	UserID      uint     `json:"user_id" gorm:"not null"`
	CategoryID  uint     `json:"category_id" gorm:"not null"`
	User        User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category    Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Reviews     []Review `json:"reviews,omitempty" gorm:"foreignKey:ProductID"`
}
