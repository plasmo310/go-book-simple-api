package services

import (
	"book-simple-api/db"
	"book-simple-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Book models.Book

type BookService struct{}

func (s BookService) GetAll() ([]Book, error) {
	var books []Book
	err := db.GetDB().Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s BookService) Create(c *gin.Context) (Book, error) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return book, err
	}
	err = db.GetDB().Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s BookService) UpdateById(id string, c *gin.Context) (Book, error) {
	var book Book
	err := db.GetDB().First(&book, id).Error
	if err != nil {
		return book, err
	}
	err = c.ShouldBindJSON(&book)
	if err != nil {
		return book, err
	}
	parseId, _ := strconv.ParseUint(id, 10, 64)
	book.Id = uint(parseId)
	err = db.GetDB().Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s BookService) DeleteById(id string, c *gin.Context) error {
	var book Book
	err := db.GetDB().First(&book, id).Error
	if err != nil {
		return err
	}
	err = db.GetDB().Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
