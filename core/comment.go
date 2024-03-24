package core

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;foreignKey" json:"user_id"`
	PhotoID   uint      `gorm:"not null;foreignKey" json:"photo_id"`
	Message   string    `gorm:"not null;type:varchar(200)" json:"message" validate:"required, max=200"`
	Photo     Photo     `json:"photo"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"-"`
}

func (Comment) TableName() string {
	return "comments"
}
