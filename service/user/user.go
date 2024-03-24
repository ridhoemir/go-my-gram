package user

import (
	"final-project/core"
	"final-project/repository/user"
)

type Service struct {
	userRepo *user.Repository
}

func NewService(userRepo *user.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) RegisterUser(user core.User) (core.User, error) {
	res, err := s.userRepo.RegisterUser(user)

	return res, err
}

func (s *Service) LoginUser(user core.User) (string, error) {
	res, err := s.userRepo.LoginUser(user)

	return res, err
}

func (s *Service) UpdateUser(user core.User) (core.User, error) {
	res, err := s.userRepo.UpdateUser(user)

	return res, err
}

func (s *Service) DeleteUser(id uint) (string, error) {
	msg, err := s.userRepo.DeleteUser(id)

	return msg, err
}
