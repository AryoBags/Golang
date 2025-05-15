package menu

import "gorm.io/gorm"

type menuRepo struct {
	db *gorm.DB
}

func Gety