package core

import (
	"time"
)

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null;type:varchar(50)" json:"name" validate:"required,min=3,max=50"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url" validate:"required, url"`
	UserID         uint      `gorm:"not null;foreignKey" json:"user_id"`
	User           User      `json:"user"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"-"`
}

func (SocialMedia) TableName() string {
	return "social_media"
}
