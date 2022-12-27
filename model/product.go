package model

import "time"

type Product struct {
	ID        uint64    `json:"id" gorm:"primaryKey;notNull"`
	Name      string    `json:"name" gorm:"notNull"`
	Price     string    `json:"price" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (Product) TableName() string {
	return "product"
}
