package photo

import "final-project/core"

type PhotoGetResponse struct {
	ID       uint              `json:"id"`
	Caption  string            `json:"caption"`
	Title    string            `json:"title"`
	PhotoUrl string            `json:"photo_url"`
	UserID   uint              `json:"user_id"`
	User     UserPhotoResponse `json:"user"`
}

type UserPhotoResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoCreateRequest struct {
	Caption  string `json:"caption" validate:"omitempty,max=200"`
	Title    string `json:"title" validate:"required,min=3,max=100"`
	PhotoUrl string `json:"photo_url" validate:"required,validateUrl"`
}

type PhotoCreateResponse struct {
	ID       uint   `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

func (p PhotoGetResponse) ToPhotoResponse(data core.Photo) PhotoGetResponse {
	return PhotoGetResponse{
		ID:       data.ID,
		Caption:  *data.Caption,
		Title:    data.Title,
		PhotoUrl: data.PhotoUrl,
		UserID:   data.UserID,
		User: UserPhotoResponse{
			ID:       data.User.ID,
			Email:    data.User.Email,
			Username: data.User.Username,
		},
	}
}

func (p PhotoCreateResponse) ToPhotoResponse(data core.Photo) PhotoCreateResponse {
	return PhotoCreateResponse{
		ID:       data.ID,
		Caption:  *data.Caption,
		Title:    data.Title,
		PhotoUrl: data.PhotoUrl,
		UserID:   data.UserID,
	}
}

func (p PhotoCreateRequest) ToPhoto(data PhotoCreateRequest) core.Photo {
	return core.Photo{
		Caption:  &data.Caption,
		Title:    data.Title,
		PhotoUrl: data.PhotoUrl,
	}
}
