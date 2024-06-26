package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response 存放返回响应

const (
	ok  = 100
	okT = 101
	okD = 102

	fail = 201
)

// Ok 成功
func Ok(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    ok,
		"info":    "success",
		"message": msg,
	})
}

// OkWithToken 成功 返回带有token的响应
func OkWithToken(ctx *gin.Context, msg string, token string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    okT,
		"info":    "success",
		"message": msg,
		"token":   token,
	})
}

// OkWithData 成功 返回带有data数据的响应
func OkWithData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    okD,
		"info":    "success",
		"message": msg,
		"data":    data,
	})
}

// Error 错误 返回带有error的响应
func Error(ctx *gin.Context, msg string, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    fail,
		"info":    "failure",
		"message": msg,
		"error":   err,
	})
}
