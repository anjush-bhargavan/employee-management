package interfaces

import (
	"context"

	"github.com/anjush-bhargavan/employee-management/pkg/models"
)

type EmployeeRepoInter interface {
	CreateEmployee(ctx context.Context, emp *models.Employee) error
	GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, emp *models.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	ListEmployees(ctx context.Context) ([]*models.Employee, error)
}
