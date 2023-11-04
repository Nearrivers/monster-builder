package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	dns := "root@tcp(127.0.0.1:3306)/monster-creator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println("gorm Db connection ", err)
		return nil, err
	}

	return db, nil
}
