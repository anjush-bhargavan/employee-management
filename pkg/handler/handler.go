package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/anjush-bhargavan/employee-management/pkg/models"
	"github.com/anjush-bhargavan/employee-management/pkg/service/interfaces"
	"github.com/anjush-bhargavan/employee-management/utility"
	"github.com/labstack/echo/v4"
)

// EmployeeHandler handles HTTP requests for employee management.
type EmployeeHandler struct {
	Service interfaces.EmployeeSvcInter
}

func NewEmployeeHandler(service interfaces.EmployeeSvcInter) *EmployeeHandler {
	return &EmployeeHandler{Service: service}
}

// CreateEmployee binds data and pass the data to service layer
// @Summary Create a new employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.CreateEmployee true "Employee Info"
// @Success 201 {object} models.Employee
// @Failure 400 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Router /employees [post]
func (h *EmployeeHandler) CreateEmployee(c echo.Context) error {

	var emp models.CreateEmployee
	if err := c.Bind(&emp); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewErrorResponse(err.Error(), "Invalid request body", http.StatusBadRequest))
	}
	employee := models.Employee{
		Name:      emp.Name,
		Position:  emp.Position,
		Salary:    emp.Salary,
		HiredDate: emp.HiredDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := h.Service.CreateEmployee(context.Background(), &employee); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewErrorResponse(err.Error(), "Could not create employee", http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, emp)
}

// GetEmployeeByID retrieve the employee details from servic layer by id
// @Summary Retrieve an employee by ID
// @Tags employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400 {object} utility.ErrorResponse
// @Failure 404 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Router /employees/{id} [get]
func (h *EmployeeHandler) GetEmployeeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewErrorResponse(err.Error(), "Invalid employee ID", http.StatusBadRequest))
	}

	emp, err := h.Service.GetEmployeeByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewErrorResponse(err.Error(), "Could not retrieve employee", http.StatusInternalServerError))
	}

	if emp == nil {
		return c.JSON(http.StatusNotFound, utility.NewErrorResponse("no data", "Employee not found", http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, emp)
}

// UpdateEmployee will update the employee details
// @Summary Update an employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body models.CreateEmployee true "Employee Info"
// @Success 200 {object} models.Employee
// @Failure 400 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Router /employees/{id} [put]
func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewErrorResponse(err.Error(), "Invalid employee ID", http.StatusBadRequest))
	}

	var emp models.CreateEmployee
	if err := c.Bind(&emp); err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewErrorResponse(err.Error(), "Invalid request body", http.StatusBadRequest))
	}
	employee := models.Employee{
		Name:      emp.Name,
		Position:  emp.Position,
		Salary:    emp.Salary,
		HiredDate: emp.HiredDate,
		UpdatedAt: time.Now(),
	}
	employee.ID = id

	if err := h.Service.UpdateEmployee(context.Background(), &employee); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewErrorResponse(err.Error(), "Could not update employee", http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, emp)
}

// DeleteEmployee delete a specific employee by ID
// @Summary Delete an employee
// @Tags employees
// @Param id path int true "Employee ID"
// @Success 204 {object} nil
// @Failure 400 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Router /employees/{id} [delete]
func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.NewErrorResponse(err.Error(), "Invalid employee ID", http.StatusBadRequest))
	}

	if err := h.Service.DeleteEmployee(context.Background(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewErrorResponse(err.Error(), "Could not delete employee", http.StatusInternalServerError))
	}

	return c.NoContent(http.StatusNoContent)
}

// ListEmployees fetch list of all employees
// @Summary List employees
// @Tags employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} utility.ErrorResponse
// @Router /employees [get]
func (h *EmployeeHandler) ListEmployees(c echo.Context) error {
	employees, err := h.Service.ListEmployees(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.NewErrorResponse(err.Error(), "Could not list employees", http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, employees)
}
