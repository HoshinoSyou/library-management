package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
	Author string `json:"author"`
	Brief  string `json:"brief"`
}

func InsertBook(name string, amount uint, author string, brief string) bool {
	dao.DB.AutoMigrate(&Book{})
	var book Book
	var b = Book{
		Model:  gorm.Model{},
		Name:   name,
		Amount: amount,
		Author: author,
		Brief:  brief,
	}
	dao.DB.Where("name = ?", name).First(&book)
	if book.ID > 0 {
		return false
	}
	dao.DB.Create(&b)
	return true
}

func QueryBookByID(id uint) (b Book) {
	dao.DB.Where("id = ?", id).First(&b)
	return
}

func QueryBookByName(name string) (b []Book) {
	dao.DB.Where("name LIKE ?", name).Find(&b)
	return
}

func QueryBookByAuthor(author string) (b []Book) {
	dao.DB.Where("author LIKE ?", author).Find(&b)
	return
}
