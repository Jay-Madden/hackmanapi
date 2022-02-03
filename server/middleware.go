package server

import (
	"github.com/gin-gonic/gin"
	"hackmanapi/data"
	"hackmanapi/data/repositories"
	"net/http"
)

func Auth(database *data.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Query("key")
		if key == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Api Key not found",
			})
			return
		}
		user, err := repositories.GetUserByKey(*database, key)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Invalid Api Key",
			})
			return
		}
		ctx.Set("User", user)

		ctx.Next()
	}
}
