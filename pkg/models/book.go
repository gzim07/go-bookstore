package models

import (
	"github.com/gzim07/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) CreateBook() *Book {
	config.Db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	config.Db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := config.Db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	config.Db.Where("ID=?", ID).Delete(&book)
	return book
}
