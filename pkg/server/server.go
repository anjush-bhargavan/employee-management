package server


import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Serverstruct struct is used to initialize the Echo Engine
type Serverstruct struct {
	E *echo.Echo
}

// StartServer is used to start the Server
func (s *Serverstruct) StartServer() {
	
	s.E.Use(middleware.Logger())
	s.E.Use(middleware.Recover())

	// Load HTML templates
	s.E.Static("/", "docs")

	// Start the server on port 8080
	s.E.Start(":8080")
}

// NewServer is used to create and initialize the server
func NewServer() *Serverstruct {
	e := echo.New()
	return &Serverstruct{
		E: e,
	}
}