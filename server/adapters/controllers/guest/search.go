package guest

import (
	gateway "api/server/adapters/gateways"
	usecase "api/server/usecases/guest"
	"github.com/labstack/echo/v4"
	"strconv"
)

type GuestController struct {
	Interactor usecase.GuestInteractor
}

func InitGuestController(sqlHandler gateway.SQLHandler) *GuestController {
	return &GuestController{
		Interactor: usecase.GuestInteractor{
			GuestPort: &gateway.GuestRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *GuestController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	guest, err := controller.Interactor.GuestByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, guest)
	return
}

func (controller *GuestController) Index(c echo.Context) (err error) {
	guests, err := controller.Interactor.Guests()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, guests)
	return
}
