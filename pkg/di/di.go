package di

import (
	"log"

	"github.com/anjush-bhargavan/employee-management/config"
	"github.com/anjush-bhargavan/employee-management/pkg/db"
	"github.com/anjush-bhargavan/employee-management/pkg/handler"
	"github.com/anjush-bhargavan/employee-management/pkg/repo"
	"github.com/anjush-bhargavan/employee-management/pkg/routes"
	"github.com/anjush-bhargavan/employee-management/pkg/server"
	"github.com/anjush-bhargavan/employee-management/pkg/service"
)

func Init() *server.Serverstruct{
	cnfg := config.LoadConfig()

	db := db.ConnectDB(*cnfg)

	redisClient, err := config.SetupRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}
	// defer db.Close()
	// defer redisClient.Client.Close()

	employeeRepo := repo.NewEmployeeRepo(db)

	employeeSvc := service.NewEmployeeService(employeeRepo,redisClient)

	employeeHandler := handler.NewEmployeeHandler(employeeSvc)

	server := server.NewServer()

	routes.InitRoutes(server.E, employeeHandler)

	

	return server
}
