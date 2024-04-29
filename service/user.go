package service

import (
	"github.com/gin-gonic/gin"
	"library-management/module"
	"library-management/util"
)

func UserLogin(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	pwd := util.Sha256(password, "usrpwd")
	return module.QueryUser(username, pwd)
}

func UserRegister(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	pwd := util.Sha256(password, "usrpwd")
	return module.InsertUser(username, pwd)
}

func UserChangePWD(ctx *gin.Context) (bool, string) {
	username := ctx.PostForm("username")
	oldPwd := ctx.PostForm("oldPassword")
	newPwd := ctx.PostForm("newPassword")
	oldp := util.Sha256(oldPwd, "usrpwd")
	newp := util.Sha256(newPwd, "usrpwd")
	return module.ChangeUserPwd(username, oldp, newp)
}
