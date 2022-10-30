package routers

import (
	"kanban-kita-api/config"
	"kanban-kita-api/controller"
	"kanban-kita-api/middleware"

	"github.com/gin-gonic/gin"
)

func RoutesList() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	handler := controller.Handlers{Connect: db}
	v1 := r.Group("/kanban/v1")
	{
		userRoutes := v1.Group("/user")
		{
			// get all users data
			userRoutes.POST("/login", handler.UserLogin)
			userRoutes.POST("/register", handler.UserRegister)
			userRoutes.PUT("/update-account", middleware.Authentication(), handler.UpdateUser)
			userRoutes.DELETE("/delete-account", middleware.Authentication(), handler.DeleteUser)
		}
		categoryRoutes := v1.Group("/categories")
		categoryRoutes.Use(middleware.Authentication(), middleware.AdminAuth())
		{
			categoryRoutes.GET("/", handler.GetAllCategory)
			categoryRoutes.POST("/", handler.CreateCategory)
			categoryRoutes.PATCH("/:id", handler.UpdateCategory)
			categoryRoutes.DELETE("/:id", handler.DeleteCtegory)
		}
		taskRoutes := v1.Group("/tasks")
		taskRoutes.Use(middleware.Authentication())
		{
			taskRoutes.GET("/", handler.GetAllTask)
			taskRoutes.POST("/", handler.CreateTask)
			taskRoutes.PUT("/:id", handler.UpdateTask)
			taskRoutes.PATCH("/:id", handler.UpdateStatusTask)
			taskRoutes.PATCH("/update-category/:id", handler.UpdateCategoryTask)
			taskRoutes.DELETE("/:id", handler.DeleteTask)
		}
	}
	return r
}
