package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// instantiate the database
	db = ConnectDatabase()

	// attach the endpoints to the controllers
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) { ctx.String(http.StatusOK, "hello, world!") })
	router.GET("/albums", getAlbums)
	router.GET("/albums/:title", getAlbumByTitle)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	// start the server
	router.Run("localhost:8000")
}
