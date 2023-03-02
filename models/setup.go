package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseUsername := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASS")
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	// databaseUsername := "root"
	// databasePassword := "123456"
	// databaseHost := "localhost"
	// databasePort := "3306"
	// databaseName := "go-todo-list"

	dsn := "" + databaseUsername + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// createDatabase := "CREATE DABASE IF NOT EXIST go-test-todolist"
	// useDatabase := "USER go-test-todolist"
	// database.Exec(createDatabase, useDatabase)

	database.AutoMigrate(&Activity{}, &Todo{})

	DB = database

}
