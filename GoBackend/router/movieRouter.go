package router

import (
	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//healthcheck endpoint
	r.GET("/movieapi/v1/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//Add movie endpoint
	r.POST("/movieapi/v1/addMovie", controller.AddMovieController)
	//Get movie endpoint
	r.GET("/movieapi/v1/getMovie/:id", controller.GetMovieController)
	//Get All movie endpoint
	r.GET("/movieapi/v1/getMovies", controller.GetAllController)
	//Update movie endpoint
	r.PUT("/movieapi/v1/updateMovie/:id", controller.UpdateController)
	//Delete movie endpoint
	r.DELETE("/movieapi/v1/deleteMovie/:id", controller.DeleteController)
	//Delete all movie endpoint
	r.DELETE("/movieapi/v1/deleteAll", controller.DeleteAllController)
	//Search movie endpoint
	r.GET("/movieapi/v1/search", controller.SearchResultController)

	return r
}
