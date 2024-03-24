package middlewares

import (
	"final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// if !checkUser(c, userData) {
		// 	return
		// }
		c.Set("userData", userData)
		c.Next()
	}
}

// func checkUser(c *gin.Context, userData interface{}) bool {
// 	db, err := database.NewDatabase()
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": "Database connection error",
// 		})
// 		return false
// 	}
// 	id := userData.(jwt.MapClaims)["id"].(float64)
// 	var user core.User
// 	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"message": "Unauthorized",
// 		})
// 		return false
// 	}
// 	return true
// }
