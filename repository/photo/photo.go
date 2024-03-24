package photo

import (
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

func (r *Repository) GetAllPhoto() ([]core.Photo, error) {
	var photo []core.Photo
	err := r.db.Debug().Preload("User").Find(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *Repository) GetPhotoById(id uint64) (core.Photo, error) {
	var photo core.Photo
	err := r.db.Debug().Preload("User").Where("id = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *Repository) CreatePhoto(photo core.Photo) (core.Photo, error) {
	err := r.db.Debug().Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *Repository) UpdatePhoto(id uint64, data core.Photo) (core.Photo, error) {
	var photo core.Photo
	err := r.db.Debug().Model(&photo).Where("id = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}

	err = r.db.Debug().Model(&photo).Where("id = ?", id).Updates(core.Photo{
		Caption:  data.Caption,
		Title:    data.Title,
		PhotoUrl: data.PhotoUrl,
	}).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *Repository) DeletePhoto(id uint64) error {
	var photo core.Photo
	err := r.db.Debug().Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return err
	}

	return nil
}
