package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"library-management/module"
	"library-management/util"
	"library-management/util/jwt"
	"log"
)

func ManagerLogin(ctx *gin.Context) (bool, string, string) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	pwd := util.Sha256(password, "mngpwd")
	res, m := module.QueryManager(username, pwd)
	if !res {
		return false, "管理员账户或密码错误！", ""
	}
	token := jwt.Create(username, m.ID, "manager")
	if token == "" {
		return false, "管理员登录状态获取失败！请联系开发者！", ""
	}
	return true, "管理员登录成功！", token
}

func ManagerRegister(ctx *gin.Context) (bool, string) {
	t, e := ctx.Get("type")
	if !e {
		return false, "用户登录状态已过期，请重新登录！"
	}
	if t != "manager" {
		return false, "没有权限！"
	}
	issuetype, exists := ctx.Get("usertype")
	if !exists {
		return false, "登录状态已过期！请重新登陆！"
	}
	if issuetype != "admin" {
		return false, "没有权限！"
	}
	uid, exists := ctx.Get("userid")
	if !exists {
		return false, "登录状态已过期！请重新登陆！"
	}

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	name := ctx.PostForm("name")
	authjson := ctx.PostForm("authorities")
	var auth []module.Authority
	err := json.Unmarshal([]byte(authjson), &auth)
	if err != nil {
		log.Printf("反序列化权限字段失败！错误信息：%v", err)
		return false, "反序列化权限字段失败！请联系开发者检查权限字段！"
	}
	res := module.InsertManager(username, util.Sha256(password, "mngpwd"), name, uid.(uint), auth)
	if !res {
		return false, "插入管理员记录失败！请联系开发者！"
	}
	return true, "创建管理员成功！"
}

func ManagerChangePwd(ctx *gin.Context) (bool, string) {
	t, e := ctx.Get("type")
	if !e {
		return false, "用户登录状态已过期，请重新登录！"
	}
	if t != "manager" {
		return false, "没有权限！"
	}
	username := ctx.PostForm("username")
	uid, exists := ctx.Get("userid")
	if !exists {
		return false, "用户登录状态已过期，请重新登录！"
	}
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	return module.ChangeManagerPwd(uid.(uint), username, util.Sha256(oldPassword, "mngpwd"), util.Sha256(newPassword, "mngpwd"))
}

func ManagerInf(ctx *gin.Context) (bool, string, module.ManagerInf) {
	t, e := ctx.Get("type")
	if !e {
		return false, "用户登录状态已过期，请重新登录！", module.ManagerInf{}
	}
	if t != "manager" {
		return false, "没有权限！", module.ManagerInf{}
	}
	uid, exists := ctx.Get("userid")
	if !exists {
		return false, "用户登录状态已过期，请重新登录！", module.ManagerInf{}
	}
	res, mf := module.QueryManagerInf(uid.(uint))
	if !res {
		return false, "找不到管理员信息！", module.ManagerInf{}
	}
	return true, "查找成功！", mf
}
