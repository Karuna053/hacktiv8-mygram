package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `json:"Name"`
	SocialMediaURL string `json:"SocialMediaURL"`
	UserID         uint   `json:"UserID"`
	User           User   `gorm:"foreignKey:user_id" json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateSocialMediaRules struct { // Used in create context.
	Name           string `validate:"required,max=255"`
	SocialMediaURL string `validate:"required,max=2000"`
}

type UpdateSocialMediaRules struct { // Used in update context.
	ID             uint   `validate:"required"`
	Name           string `validate:"required,max=255"`
	SocialMediaURL string `validate:"required,max=2000"`
}

type DeleteSocialMediaRules struct { // Used in delete context.
	ID uint `validate:"required"`
}
