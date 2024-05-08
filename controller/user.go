package controller

import (
	"github.com/gin-gonic/gin"
	"library-management/service"
	"library-management/util/response"
)

func UserLogin(ctx *gin.Context) {
	res, token := service.UserLogin(ctx)
	if res {
		response.OkWithToken(ctx, "登陆成功！", token)
	} else {
		response.Error(ctx, "用户名或密码错误！", nil)
	}
}

func UserRegister(ctx *gin.Context) {
	res, msg := service.UserRegister(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
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

func UserLogout(ctx *gin.Context) {
	res, msg := service.UserLogout(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func UserInformation(ctx *gin.Context) {
	res, msg, inf := service.InformationQuery(ctx)
	if res {
		response.OkWithData(ctx, msg, inf)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func UserInformationChange(ctx *gin.Context) {
	res, msg := service.InformationChange(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
	}
}
