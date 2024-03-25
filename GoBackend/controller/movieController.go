package controller

import (
	"fmt"
	"net/http"

	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/helper"
	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/model"
	"github.com/gin-gonic/gin"
)

func AddMovieController(c *gin.Context) {
	var movie model.Movie
	err := c.Bind(&movie)
	if err != nil {
		panic(err)
	}
	result := helper.AddMovie(movie)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{
		"Inserted Movie": movie,
	})
}

func GetMovieController(c *gin.Context) {
	id := c.Param("id")
	movie := helper.GetMovieById(id)
	c.JSON(http.StatusOK, gin.H{
		"Received Movie": movie,
	})
}

func GetAllController(c *gin.Context) {
	movies := helper.GetAllMovies()
	c.JSON(http.StatusOK, gin.H{
		"All Movies": movies,
	})
}

func UpdateController(c *gin.Context) {
	id := c.Param("id")
	var movie model.Movie
	c.Bind(&movie)
	movieResult := helper.UpdateMovie(id, movie)
	c.JSON(http.StatusOK, gin.H{
		"All Movies": movieResult,
	})
}

func DeleteController(c *gin.Context) {
	id := c.Param("id")
	result := helper.DeleteById(id)
	c.JSON(http.StatusOK, gin.H{
		"Movies": result,
	})
}

func DeleteAllController(c *gin.Context) {
	result := helper.DeleteAll()
	c.JSON(http.StatusOK, gin.H{
		"Movies": result,
	})
}

func SearchResultController(c *gin.Context) {
	keyword := c.Query("search")
	movies := helper.SearchResult(keyword)
	c.JSON(http.StatusOK, gin.H{
		"Searched Movies": movies,
	})
}
