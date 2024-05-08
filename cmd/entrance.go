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
		usr.PATCH("/changePwd", controller.UserChangePwd)
		usr.DELETE("/logout", controller.UserLogout)
		usr.GET("/information", controller.UserInformation)
		usr.PUT("/information/change", controller.UserInformationChange)
	}
	mng := routine.Group("/manager")
	{
		mng.GET("/login", controller.ManagerLogin)
		mng.Use(jwt.CheckToken())
		mng.POST("/register", controller.ManagerRegister)
		mng.PUT("/changePwd", controller.ManagerChangePwd)
		mng.GET("/information", controller.ManagerInfQuery)
	}
	book := routine.Group("/book")
	{
		book.GET("/detail", controller.Detail)
		book.Use(jwt.CheckToken())
		book.POST("/rent")
		book.POST("/return")
	}
	err := routine.Run(":8080")
	if err != nil {
		log.Printf("routine run failed! Error:%v", err)
		return
	}
}
