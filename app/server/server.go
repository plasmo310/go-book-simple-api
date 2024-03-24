package server

import (
	"book-simple-api/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := CreateRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func CreateRouter() *gin.Engine {
	r := gin.Default()

	bookController := controllers.BookController{}
	r.GET("/books", bookController.Index)
	r.POST("/books", bookController.Create)
	r.PUT("/books/:id", bookController.Update)
	r.DELETE("/books/:id", bookController.Delete)

	return r
}
