package routers

import (
	"awesomeGin/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(router *gin.Engine) {
	router.POST("/user-create", controllers.UserCreate)
	router.GET("/users", controllers.UsersShow)
	router.PUT("/user/:id", controllers.UserUpdate)
	router.GET("/user/:id", controllers.UserShow)
	router.DELETE("/user/:id", controllers.UserDelete)
	router.POST("login/", controllers.Login)
	router.PUT("/logout", controllers.Logout)

	router.GET("test/", controllers.EmployeeShow)
}
