package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoutes(e *echo.Group) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	e.GET("/films", h.FindFilm)
	e.GET("/film/:id", h.GetFilm)
	e.POST("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm)))
	e.DELETE("/film/:id", middleware.Auth(h.DeleteFilm))
	e.PATCH("/film/:id", middleware.Auth(h.UpdateFilm))
}
