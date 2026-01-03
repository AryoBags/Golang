package rest
import (
	"github.com/aryo/go-restaurant-app/internal/usecase/resto"
)
type handler struct {
	restUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restUsecase: restoUsecase,
	}
}