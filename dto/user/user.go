package user

import (
	"final-project/core"
	"time"
)

type UserCreateRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=50"`
	Password        string `json:"password" validate:"required,min=6,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Age             int    `json:"age" validate:"required,numeric,min=8"`
	ProfileImageUrl string `json:"profile_image_url" validate:"omitempty,url"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserUpdateRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Age             int    `json:"age" validate:"required,numeric,min=8"`
	ProfileImageUrl string `json:"profile_image_url" validate:"omitempty,url"`
}

func (req UserCreateRequest) ToUser() core.User {
	return core.User{
		Username:        req.Username,
		Password:        req.Password,
		Email:           req.Email,
		Age:             req.Age,
		ProfileImageUrl: &req.ProfileImageUrl,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (req UserLoginRequest) ToUser() core.User {
	return core.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req UserUpdateRequest) ToUser() core.User {
	return core.User{
		Username:        req.Username,
		Email:           req.Email,
		Age:             req.Age,
		ProfileImageUrl: &req.ProfileImageUrl,
		UpdatedAt:       time.Now(),
	}
}
