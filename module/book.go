package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Isbn   string `json:"isbn"`
	Amount uint   `json:"amount"`
	Author string `json:"author"`
	Brief  string `json:"brief"`
	Tags   []Tag  `json:"tags" gorm:"many2many:books_tags"`
}

type Rent struct {
	gorm.Model
	BookId   uint   `json:"bookId"`
	BookName string `json:"bookName"`
	UserId   uint   `json:"userId"`
}

func InsertBook(name string, isbn string, amount uint, author string, brief string, tags []Tag) bool {
	dao.DB.AutoMigrate(&Book{})
	var book Book
	var b = Book{
		Model:  gorm.Model{},
		Name:   name,
		Isbn:   isbn,
		Amount: amount,
		Author: author,
		Brief:  brief,
		Tags:   tags,
	}
	dao.DB.Table("books").Where("name = ?", name).First(&book)
	if book.ID > 0 {
		return false
	}
	dao.DB.Create(&b)
	return true
}

func QueryBookByID(id uint) (b Book) {
	dao.DB.Table("books").Where("id = ?", id).First(&b)
	return
}

func QueryBookByName(name string) (b []Book) {
	dao.DB.Table("books").Where("name LIKE ?", name).Find(&b)
	return
}

func QueryBookByAuthor(author string) (b []Book) {
	dao.DB.Table("books").Where("author LIKE ?", author).Find(&b)
	return
}
