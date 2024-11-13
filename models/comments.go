package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Comment struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `json:"user_id"`
	PhotoID   uint   `json:"photo_id"`
	Message   string `json:"message" validate:"required,max:6000"`
	User      User   `gorm:"foreignKey:user_id" json:"-"`
	Photo     Photo  `gorm:"foreignKey:photo_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
