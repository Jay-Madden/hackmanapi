package server

import (
	"github.com/gin-gonic/gin"
	"hackmanapi/data"
	"hackmanapi/data/models"
	"hackmanapi/data/repositories"
	"net/http"
	"time"
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

func RateLimit() gin.HandlerFunc {
	limits := make(map[int]time.Time)

	return func(ctx *gin.Context) {
		if val, ok := limits[ctx.Keys["User"].(models.User).Id]; !ok {
			limits[ctx.Keys["User"].(models.User).Id] = time.Now()
			return
		} else if inTimeSpan(time.Now().Add(-5*time.Second), time.Now(), val) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Message": "You are being rate limited, Please limit requests to one per 5 seconds",
			})
		}
		limits[ctx.Keys["User"].(models.User).Id] = time.Now()
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		before, after := check.Before(start), check.After(end)
		return !before && !after
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
