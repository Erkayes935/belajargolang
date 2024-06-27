package model

import "time"

type (
	Product struct {
		ID        uint64     `json:"id" gorm:"column:id;autoIncrement"`
		Name      string     `json:"name" gorm:"column:name"`
		Price     int        `json:"price" gorm:"column:price"`
		DeletedAt *time.Time `json:"-" gorm:"-"`
	}

	ProductCreate struct {
		Name  string `json:"name" gorm:"column:name"`
		Price int    `json:"price" gorm:"column:price"`
	}

	ProductUpdate struct {
		Name  string `json:"name" gorm:"column:name"`
		Price int    `json:"price" gorm:"column:price"`
	}

	Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Token struct {
		AuthToken    string `json:"auth_token"`
		SessionToken string `json:"session_token"`
	}
)
