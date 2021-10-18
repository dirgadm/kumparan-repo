package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB interact with database
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	database, err := gorm.Open("mysql", DbURL(BuildDBConfig()))

	if err != nil {
		panic("Failed connect to database!")
	}

	database.AutoMigrate(&Articles{})

	DB = database
	DB.LogMode(true)
}

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// BuildDBConfig ...
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "meligodirga",
		DBName:   "kumparan",
	}

	return &dbConfig
}

// DbURL ...
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
