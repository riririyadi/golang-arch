package database

import (
	"log"

	"github.com/riririyadi/golang-arch.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=gobackend port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Couldn't connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
