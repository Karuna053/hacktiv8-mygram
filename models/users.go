package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type User struct { // Master model. Will automigrate.
	ID          uint          `gorm:"primarykey"`
	Username    string        `json:"Username" gorm:"uniqueIndex"`
	Email       string        `json:"Email" gorm:"uniqueIndex"`
	Password    string        `json:"Password"`
	Age         int           `json:"Age"`
	Photo       []Photo       `gorm:"foreignKey:user_id"`
	SocialMedia []SocialMedia `gorm:"foreignKey:user_id"`
	Comment     []Comment     `gorm:"foreignKey:user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRegisterRules struct { // Used in register context.
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
	Age      int    `validate:"required,gt=8"`
}

type UserLoginRules struct { // Used in login context.
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
