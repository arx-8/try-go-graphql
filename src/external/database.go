package external

import (
	"github.com/arx-8/try-go-graphql/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
