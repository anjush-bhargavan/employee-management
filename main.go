package main

import (
	"github.com/anjush-bhargavan/employee-management/pkg/di"
	_ "github.com/anjush-bhargavan/employee-management/docs" 
)

// @title Employee Management API
// @version 1.0
// @description This is an API for managing employees.
// @host localhost:8080
// @BasePath /
func main() {
	server := di.Init()

	server.StartServer()
}
