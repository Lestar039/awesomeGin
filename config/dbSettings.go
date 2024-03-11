package config

import (
	"database/sql"
	_ "github.com/godror/godror"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DBConnection() {
	log.Println("initial Postgres database connect……")
	var err error
	dbUrl := os.Getenv("DB_CONNECTION_URL")
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Filed to Postgres connect to database")
	}
}

var ODB *sql.DB

func OracleDb() {
	log.Println("initial Oracle database connect……")
	var err error
	dbUrl := os.Getenv("ORACLE_URL")
	ODB, err = sql.Open("godror", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}
}
