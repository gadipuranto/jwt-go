package configs

import (
	"fmt"
	"go-jwt/domains/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	HOST       = "localhost"
	USER       = "postgres"
	DbPassword = "changeme"
	DbPort     = "5432"
	DbName     = "jwt-go"
	DBClient   *gorm.DB
	err        error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, DbPassword, DbName, DbPort)
	DBClient, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	fmt.Println("sukses koneksi database")
	DBClient.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return DBClient
}
