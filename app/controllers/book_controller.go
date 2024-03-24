package controllers

import (
	"book-simple-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct{}

func (ctrl BookController) Index(c *gin.Context) {
	var s services.BookService
	books, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"books": books})
	}
}

func (ctrl BookController) Create(c *gin.Context) {
	var s services.BookService
	book, err := s.Create(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (ctrl BookController) Update(c *gin.Context) {
	var s services.BookService
	id := c.Param("id")
	book, err := s.UpdateById(id, c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (ctrl BookController) Delete(c *gin.Context) {
	var s services.BookService
	id := c.Param("id")
	err := s.DeleteById(id, c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success deleted."})
	}
}
