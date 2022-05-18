package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 4 - Create a function called getAlbum that takes an id as a parameter and returns an Album
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

// 1 - Create a new type called Album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 2 - Create a slice of Albums
var albums = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 10.99},
	{ID: "2", Title: "The Wall", Artist: "Pink Floyd", Price: 10.99},
	{ID: "3", Title: "Wish You Were Here", Artist: "Pink Floyd", Price: 10.99},
}

// 3 - Create a function called getAlbums that returns a slice of Albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// 5 - Create a function called postAlbums that takes an Album as a parameter and adds it to the slice
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to the newAlbum variable
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the albums slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}

// 6 - Create a function called getAlbum that takes an id as a parameter and returns an Album
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
