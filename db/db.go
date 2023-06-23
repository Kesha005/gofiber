package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(url string)*gorm.DB{
	db,err := gorm.Open(mysql.Open(url),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	return db
}