package routes

import (
	"github.com/labstack/echo"
)

func SetResumenRoutes(c echo.Context) error {

	return c.JSON(202, "prueba")
}
