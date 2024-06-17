package routes

import (
	"github.com/anjush-bhargavan/employee-management/pkg/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo, h *handler.EmployeeHandler) {
	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	
	e.POST("/employees", h.CreateEmployee)
	e.GET("/employees/:id", h.GetEmployeeByID)
	e.PUT("/employees/:id", h.UpdateEmployee)
	e.DELETE("/employees/:id", h.DeleteEmployee)
	e.GET("/employees", h.ListEmployees)

}
