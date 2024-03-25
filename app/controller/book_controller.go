package controller

import (
	"book-simple-api/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct{}

func (ctrl BookController) Routes() []Route {
	return Routes{
		Route{
			MethodType:  GET,
			Path:        "/books",
			HandlerFunc: ctrl.Index,
		},
		Route{
			MethodType:  POST,
			Path:        "/books",
			HandlerFunc: ctrl.Create,
		},
		Route{
			MethodType:  PUT,
			Path:        "/books:id",
			HandlerFunc: ctrl.Update,
		},
		Route{
			MethodType:  DELETE,
			Path:        "/books:id",
			HandlerFunc: ctrl.Delete,
		},
	}
}

func (ctrl BookController) Index(c *gin.Context) {
	var r repository.BookRepository
	books, err := r.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"books": books})
	}
}

func (ctrl BookController) Create(c *gin.Context) {
	var r repository.BookRepository
	book, err := r.Create(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (ctrl BookController) Update(c *gin.Context) {
	var r repository.BookRepository
	id := c.Param("id")
	book, err := r.UpdateById(id, c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (ctrl BookController) Delete(c *gin.Context) {
	var r repository.BookRepository
	id := c.Param("id")
	err := r.DeleteById(id, c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success deleted."})
	}
}
