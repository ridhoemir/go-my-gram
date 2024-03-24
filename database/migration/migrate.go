package main

import (
	"final-project/core"
	"final-project/database"
)

func Migrate() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(core.User{}, core.SocialMedia{}, core.Photo{}, core.Comment{})
}
