package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
	UserID         uint   `json:"user_id"`
	User           User   `gorm:"foreignKey:user_id" json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
