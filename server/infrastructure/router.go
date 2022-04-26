package infrastructure

import (
	"api/server/infrastructure/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.InitGuest(e)

	e.Logger.Fatal(e.Start(":1323"))
}
