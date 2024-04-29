package controller

import (
	"github.com/gin-gonic/gin"
	"library-management/service"
	"library-management/util/response"
)

func UserLogin(ctx *gin.Context) {
	res := service.UserLogin(ctx)
	if res {
		response.Ok(ctx, "登陆成功！")
	} else {
		response.Error(ctx, "用户名或密码错误！", nil)
	}
}

func UserRegister(ctx *gin.Context) {
	res := service.UserRegister(ctx)
	if res {
		response.Ok(ctx, "注册成功！")
	} else {
		response.Error(ctx, "注册失败！请联系管理员", nil)
	}
}

func UserChangePwd(ctx *gin.Context) {
	res, msg := service.UserChangePWD(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func UserInformation(ctx *gin.Context) {

}
