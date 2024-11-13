package models

import "time"

type SocialMedia struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `gorm:"foreignKey:user_id" json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
