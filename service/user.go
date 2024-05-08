package service

import (
	"github.com/gin-gonic/gin"
	"library-management/module"
	"library-management/util"
	"library-management/util/jwt"
)

func UserLogin(ctx *gin.Context) (bool, string) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	pwd := util.Sha256(password, "usrpwd")
	res, user := module.UserQuery(username, pwd)
	token := jwt.Create(username, user.ID, "user")
	return res, token
}

func UserRegister(ctx *gin.Context) (bool, string) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	name := ctx.PostForm("name")
	pwd := util.Sha256(password, "usrpwd")
	res := module.UserInsert(username, pwd, name)
	return res, "注册失败！请联系管理员！"
}

func UserChangePWD(ctx *gin.Context) (bool, string) {
	username := ctx.PostForm("username")
	oldPwd := ctx.PostForm("oldPassword")
	newPwd := ctx.PostForm("newPassword")
	oldp := util.Sha256(oldPwd, "usrpwd")
	newp := util.Sha256(newPwd, "usrpwd")
	res, user := module.UserQuery(username, oldp)
	if !res {
		return false, "用户名或密码错误！校验失败！"
	}
	res = module.UserChangePwd(user.ID, newp)
	if !res {
		return res, "修改密码失败！请联系管理员！"
	}
	return res, "修改密码成功！"
}

func UserLogout(ctx *gin.Context) (bool, string) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	pwd := util.Sha256(password, "usrpwd")
	res, user := module.UserQuery(username, pwd)
	if !res {
		return false, "用户名或密码错误！校验失败！"
	}
	res = module.UserLogout(user.ID)
	if !res {
		return false, "注销用户失败！请联系管理员！"
	}
	return true, "注销用户成功！"
}

func InformationQuery(ctx *gin.Context) (bool, string, module.Information) {
	userid, exists := ctx.Get("userid")
	if !exists {
		return false, "登录状态已过期！请重新登陆！", module.Information{}
	}
	res, i := module.UserQueryInformation(userid.(uint))
	if !res {
		return false, "找不到用户对应信息！", module.Information{}
	}
	return res, "查找用户信息成功！", i
}

func InformationChange(ctx *gin.Context) (bool, string) {
	uid, exists := ctx.Get("userid")
	if !exists {
		return false, "登录状态已过期！请重新登陆！"
	} else {
		name := ctx.PostForm("name")
		res := module.UserChangeInformation(uid.(uint), name)
		if !res {
			return false, "修改记录失败！请联系管理员！"
		}
		return true, "修改成功！"
	}
}
