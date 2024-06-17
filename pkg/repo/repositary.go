package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/anjush-bhargavan/employee-management/pkg/models"
	"github.com/anjush-bhargavan/employee-management/pkg/repo/interfaces"
)

type employeeRepo struct {
	db *sql.DB
}

// NewEmployeeRepo creates an instance of employee repository
func NewEmployeeRepo(db *sql.DB) interfaces.EmployeeRepoInter {
	return &employeeRepo{db: db}
}

// CreateEmployee creates a new employee data in database.
func (e *employeeRepo) CreateEmployee(ctx context.Context, emp *models.Employee) error {
	query := `INSERT INTO employees (name, position, salary, hired_date, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := e.db.QueryRowContext(ctx, query, emp.Name, emp.Position, emp.Salary, emp.HiredDate, time.Now()).Scan(&emp.ID)

	return err
}

// GetEmployeeByID retrieves data if there exist an employee with given ID.
func (e *employeeRepo) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	var emp models.Employee
	query := `SELECT id, name, position, salary, hired_date, created_at, updated_at FROM employees WHERE id = $1`
	if err := e.db.QueryRowContext(ctx, query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary, &emp.HiredDate, &emp.CreatedAt, &emp.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &emp, nil
}

// UpdateEmployee updates the employee details in database by ID.
func (e *employeeRepo) UpdateEmployee(ctx context.Context, emp *models.Employee) error {
	query := `UPDATE employees SET name = $1, position = $2, salary = $3, hired_date = $4, updated_at = $5 WHERE id = $6`
	_, err := e.db.ExecContext(ctx, query, emp.Name, emp.Position, emp.Salary, emp.HiredDate, time.Now(), emp.ID)
	return err
}

// DeleteEmployee removes the employee details using ID.
func (e *employeeRepo) DeleteEmployee(ctx context.Context, id int) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := e.db.ExecContext(ctx, query, id)
	return err
}

// ListEmployees will fetch all employees
func (e *employeeRepo) ListEmployees(ctx context.Context) ([]*models.Employee, error) {
	query := `SELECT * FROM employees`
	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []*models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary, &emp.HiredDate, &emp.CreatedAt, &emp.UpdatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, &emp)
	}

	return employees, nil
}
