package middlewares

import (
	"errors"
	"final-project/core"
	"final-project/database"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		modelId, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		switch name {
		case "comment":
			CheckDBComment(c, userId, modelId)
		case "photo":
			CheckDBPhoto(c, userId, modelId)
		case "social_media":
			CheckDBSocMed(c, userId, modelId)
		default:
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid model name"})
			return
		}
	}

}

func CheckDBComment(c *gin.Context, userId uint, modelId uint64) {
	var comment core.Comment
	db := database.GetDB()

	err := db.Debug().Where("id = ?", modelId).First(&comment).Error
	if err != nil {
		err := fmt.Errorf("comment with id:%v not found", modelId)
		ErrorMsg(c, err)
		return
	}
	if comment.UserID != userId {
		err := errors.New("you are not authorized to access this data")
		ErrorMsg(c, err)
		return
	}
	c.Next()
}
func CheckDBSocMed(c *gin.Context, userId uint, modelId uint64) {
	var socMed core.SocialMedia
	db := database.GetDB()

	err := db.Debug().Where("id = ?", modelId).First(&socMed).Error

	if err != nil {
		err := fmt.Errorf("social media with id:%v not found", modelId)
		ErrorMsg(c, err)
		return
	}
	if socMed.UserID != userId {
		err := errors.New("you are not authorized to access this data")
		ErrorMsg(c, err)
		return
	}
	c.Next()
}
func CheckDBPhoto(c *gin.Context, userId uint, modelId uint64) {
	var photo core.Photo
	db := database.GetDB()

	err := db.Debug().Where("id = ?", modelId).First(&photo).Error

	if err != nil {
		err := fmt.Errorf("photo with id:%v not found", modelId)
		ErrorMsg(c, err)
		return
	}
	if photo.UserID != userId {
		err := errors.New("you are not authorized to access this data")
		ErrorMsg(c, err)
		return
	}
	c.Next()
}
func ErrorMsg(c *gin.Context, err error) {
	c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
}
