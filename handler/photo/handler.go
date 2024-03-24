package photo

import (
	photoDto "final-project/dto/photo"
	"final-project/helpers"
	"final-project/service/photo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service *photo.Service
}

func NewHandler(service *photo.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreatePhoto(ctx *gin.Context) {
	var req photoDto.PhotoCreateRequest
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

	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	data := req.ToPhoto(req)
	data.UserID = userId

	res, err := h.service.CreatePhoto(data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, res)
}

func (h *Handler) GetAllPhoto(ctx *gin.Context) {
	res, err := h.service.GetAllPhoto()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) GetPhotoById(ctx *gin.Context) {
	photoId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}
	res, err := h.service.GetPhotoById(photoId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) UpdatePhoto(ctx *gin.Context) {
	var req photoDto.PhotoCreateRequest
	photoId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}

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

	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	data := req.ToPhoto(req)
	data.UserID = userId

	res, err := h.service.UpdatePhoto(photoId, data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.CustomErrorMsg(ctx, 400, "Invalid ID")
		return
	}
	err = h.service.DeletePhoto(photoId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Photo deleted successfully"})
}
