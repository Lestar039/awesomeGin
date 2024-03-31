package controllers

import (
	"awesomeGin/config"
	"awesomeGin/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ManagerCreate(c *gin.Context) {
	var managerStruct struct {
		Number uint
		Name   string
	}
	c.Bind(&managerStruct)

	manager := models.Manager{Number: managerStruct.Number, Name: managerStruct.Name}
	result := config.DB.Create(&manager)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"manager": manager,
	})
}

func GetManager(c *gin.Context) {
	var manager models.Manager
	id := c.Param("id")
	result := config.DB.Preload("Functions").First(&manager, "id=?", id)
	fmt.Println(manager.Functions)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"manager": manager,
	})
}

func ObjectCreate(c *gin.Context) {
	var objectStruct struct {
		Number uint
		Type   string
	}
	c.Bind(&objectStruct)

	object := models.Object{Number: objectStruct.Number, Type: objectStruct.Type}
	result := config.DB.Create(&object)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"object": object,
	})
}

func FunctionCreate(c *gin.Context) {
	var functionStruct struct {
		ManagerID      uuid.UUID
		ObjectID       uuid.UUID
		Function       string
		FunctionNumber uint
		FunctionType   string
	}
	c.Bind(&functionStruct)

	var manager models.Manager
	managerResult := config.DB.Find(&manager, functionStruct.ManagerID)
	if managerResult.Error != nil {
		return
	}
	var object models.Object
	objectResult := config.DB.Find(&object, functionStruct.ObjectID)
	if objectResult.Error != nil {
		return
	}
	function := models.ManagerFunction{
		Manager:        manager,
		Object:         object,
		Function:       functionStruct.Function,
		FunctionType:   functionStruct.FunctionType,
		FunctionNumber: functionStruct.FunctionNumber,
	}
	result := config.DB.Create(&function)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"function": function,
	})
}
