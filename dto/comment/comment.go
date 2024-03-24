package comment

import "final-project/core"

type CommentCreateRequest struct {
	Message string `json:"message" validate:"required,min=1,max=200"`
	PhotoID uint   `json:"photo_id" validate:"required"`
}

type CommentCreateResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
}

type CommentGetResponse struct {
	ID      uint         `json:"id"`
	Message string       `json:"message"`
	PhotoID uint         `json:"photo_id"`
	UserID  uint         `json:"user_id"`
	User    UserComment  `json:"user"`
	Photo   PhotoComment `json:"photo"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" validate:"required,min=1,max=200"`
}

type CommentUpdateResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
}

type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

func (req CommentCreateRequest) ToComment() core.Comment {
	return core.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
	}
}

func (res CommentCreateResponse) ToCreateResponse(data core.Comment) CommentCreateResponse {
	return CommentCreateResponse{
		ID:      data.ID,
		Message: data.Message,
		PhotoID: data.PhotoID,
		UserID:  data.UserID,
	}
}

func (res CommentGetResponse) ToGetResponse(data core.Comment) CommentGetResponse {
	return CommentGetResponse{
		ID:      data.ID,
		Message: data.Message,
		PhotoID: data.PhotoID,
		UserID:  data.UserID,
		User: UserComment{
			ID:       data.User.ID,
			Email:    data.User.Email,
			Username: data.User.Username,
		},
		Photo: PhotoComment{
			ID:       data.Photo.ID,
			Caption:  *data.Photo.Caption,
			Title:    data.Photo.Title,
			PhotoUrl: data.Photo.PhotoUrl,
			UserID:   data.Photo.UserID,
		},
	}
}

func (res CommentUpdateResponse) ToUpdateResponse(data core.Comment) CommentUpdateResponse {
	return CommentUpdateResponse{
		ID:      data.ID,
		Message: data.Message,
		PhotoID: data.PhotoID,
		UserID:  data.UserID,
	}
}

func (res CommentUpdateRequest) ToComment() core.Comment {
	return core.Comment{
		Message: res.Message,
	}
}
