package models

import "time"

type Photo struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:user_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
