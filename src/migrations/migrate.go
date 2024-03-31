package main

import (
	"awesomeGin/config"
	"awesomeGin/internal/models"
	"fmt"
)

func init() {
	config.EnvReader()
	config.DBConnection()
}

func main() {
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Manager{},
		&models.Object{},
		&models.ManagerFunction{},
		&models.Relations{},
	); err != nil {
		fmt.Println("Postgres connection error")
		fmt.Println(err)
	}
}
