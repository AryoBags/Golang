package database

import (
	"github.com/aryo/go-restaurant-app/internal/model"
	"github.com/aryo/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}
	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     9500,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Jeruk",
			OrderCode: "jus_jeruk",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
	}


	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}