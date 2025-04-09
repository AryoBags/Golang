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

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Pizza",
			OrderCode: "PZ01",
			Price:     10,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Burger",
			OrderCode: "BG01",
			Price:     8,
			Type:      MenuTypeFood,
		},
	}
	drinkMenu := []MenuItem {
		{
			Name:      "Soda",
			OrderCode: "SD01",
			Price:     2,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Teh",
			OrderCode: "EH01",
			Price:     1,
			Type:      MenuTypeDrink,
		},
	}
	fmt.Println(foodMenu, drinkMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&MenuItem{})
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

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
