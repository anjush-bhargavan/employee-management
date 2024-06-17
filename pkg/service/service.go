package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/anjush-bhargavan/employee-management/config"
	"github.com/anjush-bhargavan/employee-management/pkg/models"
	inter "github.com/anjush-bhargavan/employee-management/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/employee-management/pkg/service/interfaces"
	"github.com/go-redis/redis/v8"
)

// employeeService struct is to connect  service layer with repositary layer
type employeeService struct {
	repo        inter.EmployeeRepoInter
	redisClient *config.RedisService
}

// NewEmployeeService creates an instance of service layer
func NewEmployeeService(repo inter.EmployeeRepoInter, redisClient *config.RedisService) interfaces.EmployeeSvcInter {
	return &employeeService{repo: repo,
		redisClient: redisClient}
}

// CreateEmployee passes the employee details to repo layer to create and delete the cached list of employees
func (s *employeeService) CreateEmployee(ctx context.Context, emp *models.Employee) error {
	err := s.repo.CreateEmployee(ctx, emp)
	fmt.Println("heloooooo")
	if err != nil {
		return err
	}
	_, err = s.redisClient.DelFromRedis("employees")
	if err != nil {
		return err
	}
	return nil
}

// GetEmployeeByID will check if the cache have the details else pass the id to repo layer to retrieve
// data of employee and store it in cache
func (s *employeeService) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	cacheKey := "employee_" + strconv.Itoa(id)
	cachedEmployee, err := s.redisClient.GetFromRedis(cacheKey)
	if err == redis.Nil {
		emp, err := s.repo.GetEmployeeByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if emp != nil {
			empData, _ := json.Marshal(emp)
			s.redisClient.SetDataInRedis(cacheKey, empData, time.Minute*10)
		}
		return emp, nil
	} else if err != nil {
		return nil, err
	}

	var emp models.Employee
	json.Unmarshal([]byte(cachedEmployee), &emp)
	return &emp, nil
}

// UpdateEmployee passes the data to repo layer to update and delete the data from redis
func (s *employeeService) UpdateEmployee(ctx context.Context, emp *models.Employee) error {
	err := s.repo.UpdateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	s.redisClient.DelFromRedis("employee_"+strconv.Itoa(int(emp.ID)), "employees")
	return nil
}

// DeleteEmployee passes the employee id to repo layer to delete and delete all linked data inn redis
func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	err := s.repo.DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	s.redisClient.DelFromRedis("employee_"+strconv.Itoa(id), "employees")
	return nil
}

// ListEmployees  will check the data in cache else it will get from repo layer and store it in redis
func (s *employeeService) ListEmployees(ctx context.Context) ([]*models.Employee, error) {
	cachedEmployees, err := s.redisClient.GetFromRedis("employees")
	if err == redis.Nil {
		employees, err := s.repo.ListEmployees(ctx)
		if err != nil {
			return nil, err
		}
		empData, _ := json.Marshal(employees)
		s.redisClient.SetDataInRedis("employees", empData, time.Minute*10)
		return employees, nil
	} else if err != nil {
		return nil, err
	}

	var employees []*models.Employee
	json.Unmarshal([]byte(cachedEmployees), &employees)
	return employees, nil
}
