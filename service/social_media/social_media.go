package social_media

import (
	"final-project/core"
	socMedDto "final-project/dto/social_media"
	"final-project/repository/social_media"
)

type Service struct {
	socialMediaRepo *social_media.Repository
}

func NewService(socialMediaRepo *social_media.Repository) *Service {
	return &Service{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *Service) GetAllSocialMedia(id uint) ([]map[string]interface{}, error) {
	var resData []map[string]interface{}
	var data map[string]interface{}
	res, err := s.socialMediaRepo.GetAllSocialMedia(id)
	for _, item := range res {
		data = map[string]interface{}{
			"id":               item.ID,
			"name":             item.Name,
			"social_media_url": item.SocialMediaUrl,
			"user_id":          item.UserID,
			"user": map[string]interface{}{
				"id":       item.User.ID,
				"email":    item.User.Email,
				"username": item.User.Username,
			},
		}
		resData = append(resData, data)
	}
	return resData, err
}

func (s *Service) GetSocialMediaById(userId uint, id uint64) (map[string]interface{}, error) {
	res, err := s.socialMediaRepo.GetSocialMediaById(userId, id)

	data := map[string]interface{}{
		"id":               res.ID,
		"name":             res.Name,
		"social_media_url": res.SocialMediaUrl,
		"user_id":          res.UserID,
		"user": map[string]interface{}{
			"id":       res.User.ID,
			"email":    res.User.Email,
			"username": res.User.Username,
		},
	}

	return data, err
}

func (s *Service) CreateSocialMedia(data core.SocialMedia) (socMedDto.SocMedCreateResponse, error) {
	res, err := s.socialMediaRepo.CreateSocialMedia(data)
	resDto := socMedDto.SocMedCreateResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
	}
	return resDto, err
}

func (s *Service) UpdateSocialMedia(id uint64, data core.SocialMedia) (core.SocialMedia, error) {
	res, err := s.socialMediaRepo.UpdateSocialMedia(id, data)

	return res, err
}

func (s *Service) DeleteSocialMedia(id uint64) error {
	err := s.socialMediaRepo.DeleteSocialMedia(id)

	return err
}
