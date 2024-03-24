package comment

import (
	"errors"
	commentDto "final-project/dto/comment"
	"final-project/helpers"
	"final-project/service/comment"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service *comment.Service
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateComment(ctx *gin.Context) {
	var req commentDto.CommentCreateRequest
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

	data := req.ToComment()
	data.UserID = userId

	res, err := h.service.CreateComment(data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, res)
}

func (h *Handler) GetAllComment(ctx *gin.Context) {
	res, err := h.service.GetAllComment()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) GetCommentById(ctx *gin.Context) {
	commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid comment id")
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetCommentById(commentId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (h *Handler) UpdateComment(ctx *gin.Context) {
	var req commentDto.CommentUpdateRequest
	commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
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

	// commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	// if err != nil {
	// 	err = errors.New("invalid comment id")
	// 	ctx.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// userData, _ := ctx.MustGet("userData").(jwt.MapClaims)
	// userId := uint(userData["id"].(float64))

	data := req.ToComment()

	res, err := h.service.UpdateComment(commentId, data)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) DeleteComment(ctx *gin.Context) {
	commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid comment id")
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteComment(commentId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Comment deleted successfully"})
}
