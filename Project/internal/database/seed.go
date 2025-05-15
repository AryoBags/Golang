package database

import (
	"github.com/AryoBags/go-restaurant-app/internal/model"
	"github.com/AryoBags/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)
func seedDB( db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Pizza",
			OrderCode: "PZ01",
			Price:     10,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Burger",
			OrderCode: "BG01",
			Price:     8,
			Type:      constant.MenuTypeFood,
		},
	}
	drinkMenu := []model.MenuItem{
		{
			Name:      "Soda",
			OrderCode: "SD01",
			Price:     2,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Teh",
			OrderCode: "EH01",
			Price:     1,
			Type:      constant.MenuTypeDrink,
		},
	}
	
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}