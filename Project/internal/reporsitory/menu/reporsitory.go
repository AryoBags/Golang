package menu

import (
	"github.com/AryoBags/go-restaurant-app/internal-model"
	"github.com/AryoBags/go-restaurant-app/internal/model"
)

type Reporsitory interface{
	GetMenu(menuType string )([]model.MenuItem, error)
}