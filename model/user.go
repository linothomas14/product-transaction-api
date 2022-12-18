package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;notNull"`
	Email     string    `json:"email" gorm:"unique;notNull"`
	Username  string    `json:"username" gorm:"unique;notNull"`
	Password  string    `json:"password" gorm:"notNull"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}
