package social_media

import (
	"errors"
	"final-project/core"

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

func (r *Repository) GetAllSocialMedia(id uint) ([]core.SocialMedia, error) {
	var socialMedia []core.SocialMedia
	err := r.db.Where("id = ?", id).First(&core.User{}).Error
	if err != nil {
		return socialMedia, errors.New("user not found")
	}
	err = r.db.Debug().Preload("User").Where("user_id = ?", id).Find(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *Repository) GetSocialMediaById(userId uint, id uint64) (core.SocialMedia, error) {
	var socialMedia core.SocialMedia
	err := r.db.Where("id = ?", userId).First(&core.User{}).Error
	if err != nil {
		return socialMedia, errors.New("user not found")
	}
	err = r.db.Debug().Preload("User").Where("user_id = ? AND id = ?", userId, id).First(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *Repository) CreateSocialMedia(socialMedia core.SocialMedia) (core.SocialMedia, error) {
	err := r.db.Debug().First(&socialMedia.User, socialMedia.UserID).Error
	if err != nil {
		return socialMedia, errors.New("error while creating social media")
	}

	err = r.db.Debug().Preload("User").Create(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *Repository) UpdateSocialMedia(id uint64, data core.SocialMedia) (core.SocialMedia, error) {
	var socialMedia core.SocialMedia

	err := r.db.Debug().Model(&socialMedia).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *Repository) DeleteSocialMedia(id uint64) error {
	var socialMedia core.SocialMedia
	err := r.db.Debug().Where("id = ?", id).Delete(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}
