package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON.
// Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`

}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
//main function
func main(){
	//initialise git router using gin.Default()
	router := gin.Default()
	//define the route and the associated functions
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.POST("/albums/:id", deleteAlbumID)
	router.Run("localhost:8080")
}
//This getAlbums function creates JSON from the slice of album structs, writing the JSON into the response.
//gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more
func getAlbums(c *gin.Context){
	//serializes struct into JSON
	c.IndentedJSON(http.StatusOK,albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context){
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err!= nil {
		return
	}

	//add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated,newAlbum)

}

func getAlbumByID(c *gin.Context){
	//Use Context.Param to retrieve the id path parameter from the URL.
	//  When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
	id := c.Param("id")

	for _, albumItem := range albums{
		if albumItem.ID == id {
			c.IndentedJSON(http.StatusOK, albumItem)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

func deleteAlbumID(c *gin.Context){
	//get the id from the url
	stringId := c.Param("id")
	//convert it to int
	id,_ := strconv.ParseInt(stringId, 10, 64)
	//condition if the id is more than the slice length
	if(int(id)> int(len(albums))){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Id is too high"})
		return
	}
	//condition if the id is negative
	if(int(id)< 0){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Id is too low"})
		return
	}
	//make an empty temp slice
	var tempAlbum = make([]album,len(albums)-1)
	//loop through all the items
	for i, albumItem := range albums{
		albumId,_ := strconv.ParseInt(albumItem.ID, 10, 64)
		//store the items whose id is not equal to the quert item id
		if  albumId != id {
			tempAlbum[i-1] = albumItem
		}
	}
	//at last copy the whole thing to originall slice
	albums = tempAlbum
	c.IndentedJSON(http.StatusOK, gin.H{"message":"Item Deleted"})
}

