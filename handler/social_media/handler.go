package social_media

import (
	"final-project/helpers"
	"final-project/service/social_media"
	"strconv"

	socMedDto "final-project/dto/social_media"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service *social_media.Service
}

func NewHandler(service *social_media.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllSocialMedia(ctx *gin.Context) {
	var socMedRes socMedDto.SocMedResponse
	socMedResArr := []socMedDto.SocMedResponse{}

	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)

	id := uint(userData["id"].(float64))
	res, err := h.service.GetAllSocialMedia(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, item := range res {
		data := socMedRes.ToSocMedArrResponse(item)
		socMedResArr = append(socMedResArr, data)
	}
	ctx.JSON(200, socMedResArr)
}

func (h *Handler) GetSocialMediaById(ctx *gin.Context) {
	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)

	id := uint(userData["id"].(float64))
	socMedId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}

	res, err := h.service.GetSocialMediaById(id, socMedId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) CreateSocialMedia(ctx *gin.Context) {
	var req socMedDto.SocMedCreateRequest
	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)

	id := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validate := helpers.NewValidator()
	err = validate.Validate(req)
	if err != nil {
		helpers.ReturnErrorMsg(ctx, err)
		return
	}

	data := req.ToSocialMedia(userData)
	data.UserID = id

	res, err := h.service.CreateSocialMedia(data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

func (h *Handler) UpdateSocialMedia(ctx *gin.Context) {
	var req socMedDto.SocMedUpdateRequest
	var res socMedDto.SocMedUpdateResponse
	socMedId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}

	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validate := helpers.NewValidator()
	err = validate.Validate(req)
	if err != nil {
		helpers.ReturnErrorMsg(ctx, err)
		return
	}

	data := req.ToSocialMedia(req)
	data.UserID = id
	data.ID = uint(socMedId)

	resData, err := h.service.UpdateSocialMedia(socMedId, data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res = res.ToSocMedUpdateResponse(resData)

	ctx.JSON(201, res)
}

func (h *Handler) DeleteSocialMedia(ctx *gin.Context) {
	socMedId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}

	err = h.service.DeleteSocialMedia(socMedId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Social Media Deleted"})
}
