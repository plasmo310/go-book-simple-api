package repository

import (
	"book-simple-api/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookRepository struct{}
type Book model.Book

func (r BookRepository) GetAll() ([]Book, error) {
	var books []Book
	err := Db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r BookRepository) Create(c *gin.Context) (Book, error) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return book, err
	}
	err = Db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r BookRepository) UpdateById(id string, c *gin.Context) (Book, error) {
	var book Book
	err := Db.First(&book, id).Error
	if err != nil {
		return book, err
	}
	err = c.ShouldBindJSON(&book)
	if err != nil {
		return book, err
	}
	parseId, _ := strconv.ParseUint(id, 10, 64)
	book.Id = uint(parseId)
	err = Db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r BookRepository) DeleteById(id string, c *gin.Context) error {
	var book Book
	err := Db.First(&book, id).Error
	if err != nil {
		return err
	}
	err = Db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
