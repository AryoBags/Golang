package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)
func (h *handler)GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restUsecase.GetMenu(menuType)
	if err != nil {
		fmt.Printf("Error fetching %s menu:", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),

		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})

}