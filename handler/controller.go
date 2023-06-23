package book

import "gorm.io/gorm"

type bookHandle struct {
	DB *gorm.DB
}