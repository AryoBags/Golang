package menu
import(
	"github.com/aryo/go-restaurant-app/internal/model"

)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}