package datastore

import (
	"github.com/1saifj/go_todo_tdd/src/domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{}, &model.Todo{})
}
