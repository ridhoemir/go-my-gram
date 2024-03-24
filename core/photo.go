package core

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null;type:varchar(100)" json:"title" validate:"required,min=3,max=100"`
	Caption   *string   `gorm:"type:varchar(200)" json:"caption" validate:"max=200"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url" validate:"required, url"`
	UserID    uint      `gorm:"not null;foreignKey" json:"user_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"-"`
	Comment   []Comment `gorm:"constraint:OnDelete:CASCADE" json:"comment"`
	User      User      `json:"-"`
}

func (Photo) TableName() string {
	return "photos"
}
