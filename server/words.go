package server

import (
	"github.com/gin-gonic/gin"
	"hackmanapi/data"
	"math/rand"
	"net/http"
	"strconv"
)

type Words struct {
	Db    *data.Database
	Words []string
}

func (word *Words) Get(ctx *gin.Context) {
	query := ctx.Query("length")

	// We got a request for a random word of any length
	if query == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"word": word.Words[rand.Intn(len(word.Words)-1)],
		})
		return
	}

	// We got a request for a random word of a certain length
	val, err := strconv.Atoi(query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter value",
		})
		return
	}
	if val < 4 || val > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Length Parameter out of bounds",
		})
		return
	}

	lengthWords := []string{}
	for _, word := range word.Words {
		if len(word) == val {
			lengthWords = append(lengthWords, word)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"word": lengthWords[rand.Intn(len(lengthWords)-1)],
	})
	return
}
