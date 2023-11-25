package main

import (
	"github.com/gin-gonic/gin"
)

var albums = []album{
	{ID: "1", Title: "Riot!", Artist: "Paramore", Price: 24.99},
	{ID: "2", Title: "Palomino", Artist: "Miranda Lambert", Price: 19.99},
	{ID: "3", Title: "Brothers", Artist: "The Black Keys", Price: 20.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8000")
}
