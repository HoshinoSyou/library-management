package controller

import (
	"github.com/gin-gonic/gin"
	"library-management/service"
	"library-management/util/response"
)

func Detail(ctx *gin.Context) {
	b, err := service.BookInformation(ctx)
	if err != nil {
		response.Error(ctx, err.Error(), err)
	} else {
		response.OkWithData(ctx, "find ok!", b)
	}
}
