package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type User struct {
	ID          uint          `gorm:"primarykey"`
	Username    string        `json:"username" gorm:"uniqueIndex" validate:"required"`
	Email       string        `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Password    string        `json:"password" validate:"required,min=6"`
	Age         int           `json:"age" validate:"required,gt=8"`
	Photo       []Photo       `gorm:"foreignKey:user_id"`
	SocialMedia []SocialMedia `gorm:"foreignKey:user_id"`
	Comment     []Comment     `gorm:"foreignKey:user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
