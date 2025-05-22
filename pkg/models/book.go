package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kvn-coderaz/go_practice/pkg/config"
)

var db *gorm.DB
// Book struct (Model)
type Book struct {
		gorm.Model
		Name string `gorm:"" json:"name"`
		Author string `json:"author"`
		Publication string `json:"publication"`
}

// Initialize the database connection
func init() {
		config.Connect()
		db = config.GetDB()
		db.AutoMigrate(&Book{})
	}

func (b *Book) CreateBook() *Book {
	// Create a new book record in the database
		db.NewRecord(b)
		db.Create(&b)
		return b
	}
func GetAllBooks() []Book {
	// Retrieve all book records from the database
		var Books []Book
		db.Find(&Books)
		return Books
	}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	// Retrieve a book record by its ID
		var getBook Book
		db := db.Where("ID=?", Id).Find(&getBook)
		return &getBook, db
	}
func DeleteBook(Id int64) Book {
	// Delete a book record by its ID
		var book Book
		db.Where("ID=?", Id).Delete(book)
		return book
	}
