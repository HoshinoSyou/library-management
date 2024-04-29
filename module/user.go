package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
)

type user struct {
	gorm.Model
	username string
	password string
}

func InsertUser(username string, password string) bool {
	dao.DB.AutoMigrate(&user{})
	var us user
	var u = user{
		Model:    gorm.Model{},
		username: username,
		password: password,
	}
	dao.DB.Where("username = ?", username).First(&us)
	if us.ID > 0 {
		return false
	}
	dao.DB.Create(&u)
	return true
}

func QueryUser(username string, password string) bool {
	var u user
	dao.DB.Where("username = ?", username).First(&u)
	if u.password != password {
		return false
	} else {
		return true
	}
}

func ChangeUserPwd(username string, oldPassword string, newPassword string) (bool, string) {
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
