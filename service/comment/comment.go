package comment

import (
	"final-project/core"
	commentDto "final-project/dto/comment"
	"final-project/repository/comment"
)

type Service struct {
	userRepo *comment.Repository
}

func NewService(userRepo *comment.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) CreateComment(comment core.Comment) (commentDto.CommentCreateResponse, error) {
	resData := commentDto.CommentCreateResponse{}
	res, err := s.userRepo.CreateComment(comment)

	resData = resData.ToCreateResponse(res)

	return resData, err
}

func (s *Service) GetAllComment() ([]commentDto.CommentGetResponse, error) {
	resData := []commentDto.CommentGetResponse{}
	res, err := s.userRepo.GetAllComment()

	for _, data := range res {
		resData = append(resData, commentDto.CommentGetResponse{}.ToGetResponse(data))
	}
	return resData, err
}

func (s *Service) GetCommentById(id uint64) (commentDto.CommentGetResponse, error) {
	resData := commentDto.CommentGetResponse{}
	res, err := s.userRepo.GetCommentById(id)
	if err != nil {
		return resData, err
	}
	resData = resData.ToGetResponse(res)

	return resData, err
}

func (s *Service) UpdateComment(id uint64, comment core.Comment) (commentDto.CommentUpdateResponse, error) {
	resData := commentDto.CommentUpdateResponse{}
	res, err := s.userRepo.UpdateComment(id, comment)

	resData = resData.ToUpdateResponse(res)

	return resData, err
}

func (s *Service) DeleteComment(id uint64) error {
	err := s.userRepo.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
