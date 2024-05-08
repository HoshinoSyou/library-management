package module

import (
	"errors"
	"github.com/jinzhu/gorm"
	"library-management/dao"
	"log"
)

type User struct {
	gorm.Model
	username string
	password string
}

type Information struct {
	gorm.Model
	Name string `json:"name"`
}

func UserInsert(username string, password string, name string) bool {
	dao.DB.AutoMigrate(&User{})
	var u = User{
		Model:    gorm.Model{},
		username: username,
		password: password,
	}

	var err error

	dao.DB.Begin()
	err = dao.DB.Table("users").Create(&u).Error
	if err != nil {
		dao.DB.Rollback()
		return false
	}
	dao.DB.Table("information").Create(&Information{
		Model: gorm.Model{},
		Name:  name,
	})
	dao.DB.Commit()
	return true
}

func UserQuery(username string, password string) (bool, User) {
	var u User
	dao.DB.Table("users").Where("username = ?", username).First(&u)
	if u.password != password {
		err := errors.New("cannot find record where username = " + username)
		log.Println(err)
		return false, User{}
	} else {
		return true, u
	}
}

func UserChangePwd(userid uint, newPassword string) bool {
	err := dao.DB.Table("users").Where("id = ?", userid).Update("Password", newPassword).Error
	if err != nil {
		log.Printf("change password where")
		return false
	}
	return true
}

func UserLogout(userid uint) bool {
	var err error
	dao.DB.Begin()
	err = dao.DB.Table("users").Delete("id= ?", userid).Error
	if err != nil {
		dao.DB.Rollback()
		log.Printf("删除用户 ID 为 %v 记录失败！错误信息%v", userid, err)
		return false
	}
	err = dao.DB.Table("information").Delete("id=?", userid).Error
	if err != nil {
		dao.DB.Rollback()
		log.Printf("删除用户 ID 为 %v 记录失败！错误信息%v", userid, err)
		return false
	}
	dao.DB.Commit()
	return true
}

func UserQueryInformation(id uint) (res bool, i Information) {
	dao.DB.Table("information").Where("id = ?", id).First(&i)
	if i.ID <= 0 {
		log.Printf("查询不到 ID 为 %v 的用户信息！", id)
		return false, Information{}
	}
	return true, i
}

func UserChangeInformation(id uint, name string) bool {
	i := Information{
		Model: gorm.Model{},
		Name:  name,
	}
	err := dao.DB.Table("information").Where("id = ?", id).Update(&i).Error
	if err != nil {
		log.Printf("更新 ID 为 %d 的信息失败！错误信息：%v", err)
		return false
	}
	return true
}
