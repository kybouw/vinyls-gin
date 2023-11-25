package main

import "net/http"
import "github.com/gin-gonic/gin"
import "fmt"

type album struct {
    ID  string  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price float64   `json:"price"`
}

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

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        fmt.Println("an error occured while binding json")
        return
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, album := range albums {
        if album.ID == id {
            c.IndentedJSON(http.StatusOK, album)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

