package core

import "time"

type User struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	Username        string        `gorm:"unique;type:varchar(50)" json:"username"`
	Password        string        `json:"-"`
	Email           string        `gorm:"unique;type:varchar(150)" json:"email"`
	Age             int           `json:"age" validate:"gte=0,lte=150"`
	ProfileImageUrl *string       `json:"profile_image_url" validate:"url"`
	CreatedAt       time.Time     `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time     `gorm:"type:timestamp" json:"-"`
	Comments        []Comment     `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	Photos          []Photo       `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnDelete:CASCADE" json:"-"`
}
