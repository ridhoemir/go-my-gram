package social_media

import (
	"final-project/core"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SocMedCreateRequest struct {
	Name           string `json:"name" validate:"required,min=3,max=50"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,validateUrl"`
}

type SocMedCreateResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
}

type SocMedUpdateRequest struct {
	Name           string `json:"name" validate:"required,min=3,max=50"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,validateUrl"`
}

type SocMedUpdateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocMedResponse struct {
	ID             uint               `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"social_media_url"`
	UserID         uint               `json:"user_id"`
	User           UserSocMedResponse `json:"user"`
}

type UserSocMedResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (req SocMedCreateRequest) ToSocialMedia(user jwt.MapClaims) core.SocialMedia {
	return core.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         uint(user["id"].(float64)),
		User: core.User{
			ID:       uint(user["id"].(float64)),
			Email:    user["email"].(string),
			Username: user["username"].(string),
		},
	}
}

func (res SocMedResponse) ToSocMedArrResponse(data map[string]interface{}) SocMedResponse {
	return SocMedResponse{
		ID:             data["id"].(uint),
		Name:           data["name"].(string),
		SocialMediaUrl: data["social_media_url"].(string),
		UserID:         data["user_id"].(uint),
		User: UserSocMedResponse{
			ID:       data["user"].(map[string]interface{})["id"].(uint),
			Email:    data["user"].(map[string]interface{})["email"].(string),
			Username: data["user"].(map[string]interface{})["username"].(string),
		},
	}
}

func (res SocMedUpdateResponse) ToSocMedUpdateResponse(data core.SocialMedia) SocMedUpdateResponse {
	return SocMedUpdateResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UserID:         data.UserID,
		UpdatedAt:      data.UpdatedAt,
	}
}

func (res SocMedUpdateRequest) ToSocialMedia(data SocMedUpdateRequest) core.SocialMedia {
	return core.SocialMedia{
		Name:           data.Name,
		SocialMediaUrl: data.SocialMediaUrl,
		UpdatedAt:      time.Now(),
	}
}
