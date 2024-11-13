package models

import "time"

type User struct {
	ID          uint          `gorm:"primarykey"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Age         int           `json:"age"`
	Photo       []Photo       `gorm:"foreignKey:user_id"`
	SocialMedia []SocialMedia `gorm:"foreignKey:user_id"`
	Comment     []Comment     `gorm:"foreignKey:user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
