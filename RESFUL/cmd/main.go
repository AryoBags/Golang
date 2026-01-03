package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost user=postgres password=postgres dbname=go_resto_app port=5432 sslmode=disable "
)
func main() {
	seedDB()
	e := echo.New()
	// localhost:14045/menu/food
	
	// localhost:14045/menu/drink
	e.Logger.Fatal(e.Start(":14045"))
}






// func getFoodMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }
// func getDrinkMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
	
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

