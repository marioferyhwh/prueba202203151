package routes

import "github.com/labstack/echo"

func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/"
	const prefixResumen = prefixAPI + "resumen"
	e.GET(prefixResumen+"/:id", SetResumenRoutes)
}
