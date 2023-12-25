package routes

import (
	"mpc-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/users", controller.CreateUserController)
	e.GET("/users", controller.GetAllUserController)
	e.GET("/users/:user_id", controller.GetUserByIdController)
	e.PUT("/users/:user_id", controller.UpdateUserByIdController)
	e.DELETE("/users/:user_id", controller.DeleteUserByIdController)
	e.GET("/users/:user_id/products", controller.GetUserProductController)
	e.POST("/products", controller.CreateProductController)
	e.PUT("/products/:product_id", controller.UpdateProductByIdController)
	e.DELETE("/products/:product_id", controller.DeleteProductByIdController)
	e.GET("/products", controller.GetAllProductController)
	e.GET("/products/:product_id", controller.GetProductByIdController)
}