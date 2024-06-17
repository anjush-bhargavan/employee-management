package db

import (
	"database/sql"
	"fmt"
	"log"
	 _ "github.com/lib/pq"
	"github.com/anjush-bhargavan/employee-management/config"
)

// ConnectDB establishes connection to database and migrate tables
func ConnectDB(config config.Config) *sql.DB {
	// Create the database connection string
	dsn := "postgresql://" + config.DBUser + ":" + config.DBPassword + "@" + config.DBHost + "/" + config.DBDatabase + "?sslmode=disable"

	// Open the database connection using the pq driver
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}

func migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		position VARCHAR(100) NOT NULL,
		salary INTEGER NOT NULL,
		hired_date DATE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to execute migration: %v", err)
	}

	fmt.Println("Database migration completed successfully!")
	return nil
}
