package database

import (
	"fmt"
	"main/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

const (
	host     = "db-service"
	port     = 5432
	user     = "postgres"
	password = "172754"
	dbname   = "postgres"
)

func Data() *gorm.DB {

	psq := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ := gorm.Open(postgres.Open(psq), &gorm.Config{})

	db.AutoMigrate(&model.Users{}, &model.Todos{}, &model.Companies{})
	return db
}
