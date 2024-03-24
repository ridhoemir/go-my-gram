package database

import (
	"final-project/core"
	"final-project/helpers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func NewDatabase() (*gorm.DB, error) {
	var (
		host     = helpers.GetEnv("DB_HOST")
		port     = helpers.GetEnv("DB_PORT")
		user     = helpers.GetEnv("DB_USERNAME")
		password = helpers.GetEnv("DB_PASSWORD")
		dbname   = helpers.GetEnv("DB_NAME")
	)
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	db.Debug().AutoMigrate(core.User{}, core.SocialMedia{}, core.Photo{}, core.Comment{})
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
