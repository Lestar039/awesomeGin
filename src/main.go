package main

import (
	"awesomeGin/config"
	"awesomeGin/internal/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.EnvReader()
	config.DBConnection()
	config.OracleDb()
}

func main() {
	router := gin.Default()
	routers.InitializeUserRoutes(router)
	router.Run()
}
