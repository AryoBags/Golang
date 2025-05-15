package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=user dbname= go_resto_app sslmode=disable"
)

func main() {

	seedDB()
	e := echo.New()
	//localhost:14045/menu/food
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}






func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

// func getFoodMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
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
// 		panic(err)
// 	}

// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }
