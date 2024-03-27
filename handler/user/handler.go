package user

import (
	userDto "final-project/dto/user"
	"final-project/helpers"
	"final-project/service/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service *user.Service
}

func NewHandler(service *user.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var req userDto.UserCreateRequest

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
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	req.Password = hashedPassword
	data := req.ToUser()

	res, err := h.service.RegisterUser(data)
	if err != nil {
		if err.Error() == "email already registered" || err.Error() == "username already registered" {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	var req userDto.UserLoginRequest

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
	data := req.ToUser()

	res, err := h.service.LoginUser(data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": res})
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var req userDto.UserUpdateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"errors": err.Error(),
			"req":    req,
		})
		return
	}
	validate := helpers.NewValidator()

	err = validate.Validate(req)
	if err != nil {
		helpers.ReturnErrorMsg(ctx, err)
		return
	}

	data := req.ToUser()
	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	data.ID = uint(userData["id"].(float64))
	res, err := h.service.UpdateUser(data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))

	res, err := h.service.DeleteUser(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": res})
}
