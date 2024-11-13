package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `json:"user_id"`
	PhotoID   uint   `json:"photo_id"`
	Message   string `json:"message"`
	User      User   `gorm:"foreignKey:user_id" json:"-"`
	Photo     Photo  `gorm:"foreignKey:photo_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
