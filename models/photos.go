package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Photo struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url" validate:"required"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:user_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
