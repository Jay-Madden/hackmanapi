package server

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"hackmanapi/data"
	"hackmanapi/data/models"
	"hackmanapi/data/repositories"
	"net/http"
	"time"
)

func Auth(database *data.Database) gin.HandlerFunc {
	userCache := cache.New(20*time.Minute, 40*time.Minute)

	return func(ctx *gin.Context) {
		key := ctx.Query("key")
		if key == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Api Key not found",
			})
			return
		}

		var (
			user   interface{}
			status bool
		)

		if user, status = userCache.Get(key); !status {
			var err error
			user, err = repositories.GetUserByKey(*database, ctx.Request.Context(), key)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Invalid Api Key",
				})
				return
			}

			userCache.Set(key, user, cache.DefaultExpiration)
		}

		ctx.Set("User", user.(models.User))

		ctx.Next()
	}
}

func RateLimitIp() gin.HandlerFunc {
	limits := make(map[string]time.Time)

	return func(ctx *gin.Context) {
		if val, ok := limits[ctx.ClientIP()]; !ok {
			limits[ctx.ClientIP()] = time.Now()
			return
		} else if inTimeSpan(time.Now().Add(-2*time.Second), time.Now(), val) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Message": "You are being rate limited, Please limit requests to one per 2 seconds",
				"Bucket":  "Ip Limit",
			})
		}
		limits[ctx.ClientIP()] = time.Now()
	}
}

func RateLimitKey() gin.HandlerFunc {
	limits := make(map[int]time.Time)

	return func(ctx *gin.Context) {
		if val, ok := limits[ctx.Keys["User"].(models.User).Id]; !ok {
			limits[ctx.Keys["User"].(models.User).Id] = time.Now()
			return
		} else if inTimeSpan(time.Now().Add(-2*time.Second), time.Now(), val) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Message": "You are being rate limited, Please limit requests to one per 2 seconds",
				"Bucket":  "Api Key",
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
