package db

import (
	"fmt"
	"nearrivers/monster-creator/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

func ConnectToDb() error {
	dns := "root@tcp(127.0.0.1:3306)/monster-creator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println("gorm Db connection ", err)
		return err
	}

	db.AutoMigrate(
		&models.Campaign{},
		&models.Monster{},
		&models.SpecialTrait{},
		&models.Action{},
	)

	dB = db

	return nil
}

func GetDbConnection() *gorm.DB {
	return dB
}
