package routes

import (
	"strconv"

	"github.com/marioferyhwh/prueba202203151/pkg/controllers"

	"github.com/labstack/echo"
)

func SetResumenRoutes(c echo.Context) error {
	fecha := c.Param("id")
	s := c.Request().URL.Query().Get("dias")

	dias, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return c.JSON(400, "dias no validos")
	}
	answer := controllers.GetResumen(fecha, dias)
	return c.JSON(202, answer)
}
