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
			Name:      "Pizza",
			OrderCode: "PZ01",
			Price:     10,
		},
		{
			Name:      "Burger",
			OrderCode: "BG01",
			Price:     8,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Soda",
			OrderCode: "SD01",
			Price:     2,
		},
		{
			Name:      "Es Teh",
			OrderCode: "EH01",
			Price:     1,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}
func main() {
	e := echo.New()
	//localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
