package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Comment struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `json:"UserID"`
	PhotoID   uint   `json:"PhotoID"`
	Message   string `json:"Message"`
	User      User   `gorm:"foreignKey:user_id" json:"-"`
	Photo     Photo  `gorm:"foreignKey:photo_id" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCommentRules struct { // Used in create context.
	PhotoID uint   `validate:"required"`
	Message string `validate:"required,max=2000"`
}

type UpdateCommentRules struct { // Used in update context.
	ID      uint   `validate:"required"`
	PhotoID uint   `validate:"required"`
	Message string `validate:"required,max=2000"`
}

type DeleteCommentRules struct { // Used in delete context.
	ID uint `validate:"required"`
}
