package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

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

type CreatePhotoRules struct { // Used in create context.
	Title    string `validate:"required,max=255"`
	Caption  string `validate:"max=2000"`
	PhotoURL string `validate:"required,max=2000"`
	UserID   string `validate:"required"`
}

type UpdatePhotoRules struct { // Used in update context.
	ID       uint   `validate:"required"`
	Title    string `validate:"required,max=255"`
	Caption  string `validate:"max=2000"`
	PhotoURL string `validate:"required,max=2000"`
	UserID   string `validate:"required"`
}

type DeletePhotoRules struct { // used in delete context.
	ID uint `validate:"required"`
}
