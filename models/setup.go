package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	databaseUsername := os.Getenv("MYSQL_USER")
	databasePassword := os.Getenv("MYSQL_PASSWORD")
	databaseHost := os.Getenv("MYSQL_HOST")
	databasePort := os.Getenv("MYSQL_PORT")
	databaseName := os.Getenv("MYSQL_DBNAME")
	// databaseUsername := "root"
	// databasePassword := "123456"
	// databaseHost := "127.0.0.1"
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
