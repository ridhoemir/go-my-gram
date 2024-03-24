package comment

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

func (r *Repository) CreateComment(comment core.Comment) (core.Comment, error) {
	err := r.db.Debug().Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *Repository) GetAllComment() ([]core.Comment, error) {
	var comment []core.Comment
	err := r.db.Debug().Preload("User").Preload("Photo").Find(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *Repository) GetCommentById(id uint64) (core.Comment, error) {
	var comment core.Comment
	err := r.db.Debug().Preload("User").Preload("Photo").Where("id = ?", id).First(&comment).Error
	if err != nil {
		err = errors.New("comment not found")
		return comment, err
	}

	return comment, nil
}

func (r *Repository) UpdateComment(id uint64, data core.Comment) (core.Comment, error) {
	var comment core.Comment
	err := r.db.Debug().Model(&comment).Where("id = ?", id).First(&comment).Error
	if err != nil {
		err = errors.New("comment not found")
		return comment, err
	}
	err = r.db.Debug().Model(&comment).Where("id = ?", id).Updates(core.Comment{
		Message: data.Message,
	}).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *Repository) DeleteComment(id uint64) error {
	var comment core.Comment
	err := r.db.Debug().Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return err
	}

	return nil
}
