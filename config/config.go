package config

import (
	"github.com/jinzhu/gorm"
	"Piruma/model"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "admin:OGTIRKWUIGESPYPB@tcp(sl-aus-syd-1-portal.5.dblayer.com:20533)/piruma")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(model.Mahasiswa{},model.History{})
	return db
}

