package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	e.GET("/categories", h.FindCategory)
	e.GET("/category/:id", h.GetCategory)
	e.POST("/category", middleware.Auth(h.CreateCategory))
	e.DELETE("/category/:id", middleware.Auth(h.DeleteCategory))
	e.PATCH("/category/:id", middleware.Auth(h.UpdateCategory))
	// }
}
