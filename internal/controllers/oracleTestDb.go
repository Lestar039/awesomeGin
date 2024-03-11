package controllers

import (
	"awesomeGin/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmployeeShow(c *gin.Context) {
	rows, err := config.ODB.Query("select * from employee")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var employee string
	for rows.Next() {

		rows.Scan(&employee)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": employee,
	})
}
