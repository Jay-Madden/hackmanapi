package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hackmanapi/data"
	"hackmanapi/server"
	"log"
	"os"
	"strings"
)

// @title HackMan Api
// @version 1.0
// @description HackMan Api to get a random word

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:4321
// @BasePath /api
// @schemes http
func main() {
	log.Println("Initializing HackManApi")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := data.InitializeDb(os.Getenv("DB_CONN_STR"))
	db.CreateTables()

	dat, err := os.ReadFile("words.txt")

	lines := strings.Split(string(dat), "\n")

	wordController := server.Words{Db: db, Words: lines}

	ginServer := gin.Default()

	ginServer.Use(server.Auth(db))
	ginServer.Use(server.RateLimit())

	ginServer.GET("/api/word", wordController.Get)

	err = ginServer.Run()
	if err != nil {
		return
	}

	db.Close()
}
