package photo

import (
	"final-project/core"
	photoDto "final-project/dto/photo"
	"final-project/repository/photo"
)

type Service struct {
	photoRepo *photo.Repository
}

func NewService(photoRepo *photo.Repository) *Service {
	return &Service{
		photoRepo: photoRepo,
	}
}

func (s *Service) GetAllPhoto() ([]photoDto.PhotoGetResponse, error) {
	photo, err := s.photoRepo.GetAllPhoto()
	arrData := []photoDto.PhotoGetResponse{}
	data := photoDto.PhotoGetResponse{}
	if err != nil {
		return arrData, err
	}
	for _, v := range photo {
		arrData = append(arrData, data.ToPhotoResponse(v))
	}
	return arrData, nil
}

func (s *Service) GetPhotoById(id uint64) (photoDto.PhotoGetResponse, error) {
	photo, err := s.photoRepo.GetPhotoById(id)
	if err != nil {
		return photoDto.PhotoGetResponse{}, err
	}
	return photoDto.PhotoGetResponse{}.ToPhotoResponse(photo), nil

}

func (s *Service) CreatePhoto(photo core.Photo) (photoDto.PhotoCreateResponse, error) {
	res, err := s.photoRepo.CreatePhoto(photo)
	if err != nil {
		return photoDto.PhotoCreateResponse{}, err
	}
	return photoDto.PhotoCreateResponse{}.ToPhotoResponse(res), nil
}

func (s *Service) UpdatePhoto(id uint64, data core.Photo) (photoDto.PhotoCreateResponse, error) {
	res, err := s.photoRepo.UpdatePhoto(id, data)
	if err != nil {
		return photoDto.PhotoCreateResponse{}, err
	}
	return photoDto.PhotoCreateResponse{}.ToPhotoResponse(res), nil
}

func (s *Service) DeletePhoto(id uint64) error {
	err := s.photoRepo.DeletePhoto(id)
	if err != nil {
		return err
	}
	return nil
}
