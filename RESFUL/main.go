package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func getDrinkMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     9500,
		},
		{
			Name:      "Jus Jeruk",
			OrderCode: "jus_jeruk",
			Price:     7000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func main() {
	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	// localhost:14045/menu/drink
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
