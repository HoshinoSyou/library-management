package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"library-management/module"
	"log"
)

func BookInformation(ctx *gin.Context) (b module.Book, err error) {
	id, exists := ctx.Get("id")
	if exists {
		b = module.QueryBookByID(id.(uint))
	} else {
		err = errors.New("cannot find bookId")
		log.Println(err)
	}
	return
}
