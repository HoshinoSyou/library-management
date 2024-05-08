package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
)

type Tag struct {
	gorm.Model
	TagName string `json:"tagName"`
}

func InsertTag(tagName string) bool {
	var t Tag
	dao.DB.Where("tag_name = ?", tagName).First(&t)
	if t.ID > 0 {
		return false
	}
	err := dao.DB.Create(&Tag{
		Model:   gorm.Model{},
		TagName: tagName,
	}).Error
	if err != nil {
		return false
	}
	return true
}

func TagList() (tags []Tag) {
	dao.DB.Table("tags").Find(&tags)
	return
}

func QueryTagWithId(tagId uint) (t Tag) {
	dao.DB.Table("tags").Where("id = ?", tagId).First(&t)
	return
}

func QueryTagWithTagName(tagName string) (t Tag) {
	dao.DB.Table("tags").Where("tagName = ?", tagName).First(&t)
	return
}

func QueryTagWithBook(bookId uint) []Tag {
	var tags []Tag
	dao.DB.Table("book").Where("book_id = ?", bookId).Find(&tags)
	return tags
}
