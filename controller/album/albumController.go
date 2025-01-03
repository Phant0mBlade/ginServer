package albumcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	albumschema "ginServer/models/album"
)

// albums slice to seed record album data.
var albums = []albumschema.Album{
	{ID: "1", FirstName: "Blue Train", LastName: "John Coltrane"},
	{ID: "2", FirstName: "Jeru", LastName: "Gerry Mulligan"},
	{ID: "3", FirstName: "Sarah Vaughan and Clifford Brown", LastName: "Sarah Vaughan"},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum albumschema.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
