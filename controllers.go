package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	var albums []album
	db.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	db.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// XXX this function no longer used
// album titles can have spaces, so pulling from the slug isn't ideal
// func getAlbumByTitle(c *gin.Context) {
// 	title := c.Param("title")
// 	var albums []album
// 	result := db.Where("title = ?", title).Find(&albums)
// 	if result.RowsAffected > 0 {
// 		c.IndentedJSON(http.StatusOK, albums)
// 		return
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var a album
	result := db.Where("id = ?", id).First(&a)
	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusOK, a)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
