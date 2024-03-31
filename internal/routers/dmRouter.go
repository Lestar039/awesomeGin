package routers

import (
	"awesomeGin/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeDMRoutes(router *gin.Engine) {
	router.POST("/manager-create", controllers.ManagerCreate)
	router.POST("/object-create", controllers.ObjectCreate)
	router.POST("/function-create", controllers.FunctionCreate)
	router.GET("/managers/:id", controllers.GetManager)
}
