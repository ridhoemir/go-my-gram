package user

import (
	"errors"
	"final-project/core"
	"final-project/helpers"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) RegisterUser(user core.User) (core.User, error) {
	result := r.db.Debug().Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected > 0 {
		err := errors.New("email already registered")
		return user, err
	}

	result = r.db.Debug().Where("username = ?", user.Username).Find(&user)
	if result.RowsAffected > 0 {
		err := errors.New("username already registered")
		return user, err
	}

	err := r.db.Debug().Create(&user).Error

	return user, err
}

func (r *Repository) LoginUser(user core.User) (string, error) {
	var token string
	var fetchedData core.User
	err := r.db.Debug().First(&fetchedData, "email = ?", user.Email).Error
	if err != nil {
		err := errors.New("wrong email/password")
		return token, err
	}

	isValid := helpers.ValidateHashPassword(user.Password, fetchedData.Password)
	if !isValid {
		err := errors.New("wrong email/password")
		return token, err
	}

	jwt, err := helpers.GenerateToken(fetchedData.ID, fetchedData.Email, fetchedData.Username)
	if err != nil {
		err = errors.New("failed to generate token")
		return token, err
	}

	return jwt, err
}

func (r *Repository) UpdateUser(user core.User) (core.User, error) {
	var fetchedData core.User
	result := r.db.Debug().Where("email = ?", user.Email).Not("id = ?", user.ID).Find(&fetchedData)
	if result.RowsAffected > 0 {
		err := errors.New("email already registered, use another email")
		return user, err
	}

	result = r.db.Debug().Where("username = ?", user.Username).Not("id = ?", user.ID).Find(&fetchedData)
	if result.RowsAffected > 0 {
		err := errors.New("username already registered, use another username")
		return user, err
	}
	err := r.db.Debug().Model(&user).Updates(&user).Error
	if err != nil {
		err = errors.New("update user failed")
		return user, err
	}

	return user, err
}

func (r *Repository) DeleteUser(id uint) (string, error) {
	var user core.User
	var msg string

	err := r.db.Debug().First(&user, "id = ?", id).Error
	if err != nil {
		err = errors.New("user not found")
		return msg, err
	}

	err = r.db.Debug().Where("id = ?", id).Delete(&user).Error
	if err != nil {
		err = errors.New("delete user failed")
		return msg, err
	}

	msg = "user deleted successfully"
	return msg, err
}
