package dao

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// Init 初始化mysql数据库
func init() {
	open, err := gorm.Open("mysql", "root:syouZX@tcp(localhost)/guguyun?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("sql init failured:%v", err)
		return
	}
	DB = open
	return
}
