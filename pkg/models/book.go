package models

import (
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/deepam02/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"` // Fixed the JSON tag syntax
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b) // Use db.Create with &b to properly create the book
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
