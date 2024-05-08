package module

import (
	"github.com/jinzhu/gorm"
	"library-management/dao"
	"log"
)

type manager struct {
	gorm.Model
	username string
	password string
}

type ManagerInf struct {
	gorm.Model
	Name        string      `json:"name"`
	Issue       uint        `json:"issue"`
	Authorities []Authority `json:"authorities" gorm:"many2many:managers_authorities"`
}

type Authority struct {
	gorm.Model
	Level string `json:"level"`
}

func InsertManager(username string, password string, name string, issue uint, auth []Authority) bool {
	dao.DB.AutoMigrate(&manager{})
	dao.DB.AutoMigrate(&ManagerInf{})
	var ma manager
	var m = manager{
		Model:    gorm.Model{},
		username: username,
		password: password,
	}
	dao.DB.Begin()
	dao.DB.Table("managers").Where("username = ?", username).First(&ma)
	if ma.ID > 0 {
		return false
	}
	err := dao.DB.Table("managers").Create(&m).Error
	if err != nil {
		dao.DB.Rollback()
		log.Printf("创建名为 %s 的管理员失败！错误信息：%v", name, err)
		return false
	}
	err = dao.DB.Table("manager_inf").Create(&ManagerInf{
		Model:       gorm.Model{},
		Name:        name,
		Issue:       issue,
		Authorities: auth,
	}).Error
	if err != nil {
		dao.DB.Rollback()
		log.Printf("创建名为 %s 的管理员失败！错误信息：%v", name, err)
		return false
	}
	return true
}

func QueryManager(username string, password string) (bool, ManagerInf) {
	var m manager
	var mf ManagerInf
	dao.DB.Table("managers").Where("username = ?", username).First(&m)
	if m.password != password {
		return false, mf
	} else {
		dao.DB.Table("manager_inf").Where("id = ?", m.ID).First(&mf)
		return true, mf
	}
}

func ChangeManagerPwd(userid uint, username string, oldPassword string, newPassword string) (bool, string) {
	var m manager
	err := dao.DB.Table("managers").Where("id = ?", userid).First(&m).Error
	if err != nil || m.ID <= 0 {
		return false, "管理员账户不存在！"
	}
	if m.username != username || m.password != oldPassword {
		return false, "管理员账户或旧密码错误！"
	}
	dao.DB.Model(&m).Table("managers").Where("id = ?", m.ID).Update("Password", newPassword)
	return true, "修改密码成功！"
}

func QueryManagerInf(userid uint) (res bool, mf ManagerInf) {
	err := dao.DB.Table("manager_inf").Where("id = ?").First(&mf).Error
	if err != nil {
		return
	}
	return true, mf
}
