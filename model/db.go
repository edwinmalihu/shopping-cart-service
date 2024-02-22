package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConnection
func DBConnection() (*gorm.DB, error) {
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	SSLMODE := os.Getenv("SSLMODE")
	TIMEZONE := os.Getenv("TIMEZONE")
	// USER := "postgres"
	// PASS := "edwin"
	// HOST := "localhost"
	// PORT := "5432"
	// DBNAME := "synapsis"
	// SSLMODE := "disable"
	// TIMEZONE := "Asia/Jakarta"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASS, DBNAME, PORT, SSLMODE, TIMEZONE)

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})

}
