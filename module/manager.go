package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
)

type manager struct {
	gorm.Model
	username string
	password string
}

func InsertManager(username string, password string) bool {
	dao.DB.AutoMigrate(&manager{})
	var ma manager
	var m = manager{
		Model:    gorm.Model{},
		username: username,
		password: password,
	}
	dao.DB.Where("username = ?", username).First(&ma)
	if ma.ID > 0 {
		return false
	}
	dao.DB.Create(&m)
	return true
}

func QueryManager(username string, password string) bool {
	var u user
	dao.DB.Where("username = ?", username).First(&u)
	if u.password != password {
		return false
	} else {
		return true
	}
}

func ChangeManagerPwd(username string, oldPassword string, newPassword string) (bool, string) {
	var u = user{
		username: username,
		password: oldPassword,
	}
	res := dao.DB.NewRecord(u)
	if !res {
		return false, "用户名或旧密码错误！"
	}
	dao.DB.Model(&u).Where("Username = ?").Update("Password", newPassword)
	return true, "修改密码成功！"
}
