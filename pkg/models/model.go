package models

import "time"

// Employee represents an employee details.
type Employee struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John"`
	Position  string    `json:"position" example:"Software Engineer"`
	Salary    int       `json:"salary" example:"60000"`
	HiredDate string    `json:"hired_date" example:"2024-06-01"`
	CreatedAt time.Time `json:"created_at" example:"2024-06-10T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-06-10T12:00:00Z"`
}

// CreateEmployee represents the details for creating a new employee.
type CreateEmployee struct {
	Name      string `json:"name" example:"John Doe"`
	Position  string `json:"position" example:"Software Engineer"`
	Salary    int    `json:"salary" example:"60000"`
	HiredDate string `json:"hired_date" example:"2024-06-01"`
}
