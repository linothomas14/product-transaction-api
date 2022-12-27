package model

import "time"

type Product struct {
	ID        uint32    `json:"id" gorm:"primaryKey;notNull"`
	Name      string    `json:"name" gorm:"notNull"`
	Price     uint32    `json:"price" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (Product) TableName() string {
	return "product"
}
