package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	FilmRoutes(e)
	CategoryRoutes(e)
	EpisodeRoutes(e)
	TransactionRoutes(e)

}
