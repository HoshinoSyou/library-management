package cmd

import (
	"github.com/gin-gonic/gin"
	"library-management/controller"
	"library-management/util/jwt"
	"log"
)

func Entrance() {
	routine := gin.Default()
	routine.GET("/home")
	usr := routine.Group("/user")
	{
		usr.GET("/login", controller.UserLogin)
		usr.POST("/register", controller.UserRegister)
		usr.Use(jwt.CheckToken())
		usr.PUT("/changePwd", controller.UserChangePwd)
		usr.GET("/information")
	}
	mng := routine.Group("/manager")
	{
		mng.GET("/login")
		mng.POST("/register")
		mng.Use(jwt.CheckToken())
		mng.PUT("/changePwd")
		mng.GET("/information")
	}
	book := routine.Group("/book")
	{
		book.GET("/detail")
		book.POST("/rent")
		book.POST("/return")
	}
	err := routine.Run(":8080")
	if err != nil {
		log.Printf("routine run failed! Error:%v", err)
		return
	}
}
