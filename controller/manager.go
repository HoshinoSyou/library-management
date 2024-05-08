package controller

import (
	"github.com/gin-gonic/gin"
	"library-management/service"
	"library-management/util/response"
)

func ManagerLogin(ctx *gin.Context) {
	res, msg, token := service.ManagerLogin(ctx)
	if res {
		response.OkWithToken(ctx, msg, token)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func ManagerRegister(ctx *gin.Context) {
	res, msg := service.ManagerRegister(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func ManagerChangePwd(ctx *gin.Context) {
	res, msg := service.ManagerChangePwd(ctx)
	if res {
		response.Ok(ctx, msg)
	} else {
		response.Error(ctx, msg, nil)
	}
}

func ManagerInfQuery(ctx *gin.Context) {
	res, msg, mf := service.ManagerInf(ctx)
	if res {
		response.OkWithData(ctx, msg, mf)
	} else {
		response.Error(ctx, msg, nil)
	}
}
