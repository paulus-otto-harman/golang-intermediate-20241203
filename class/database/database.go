package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"project/class/config"
	"time"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	// Configure the database logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color output
		},
	)

	// Open the connection to the database
	db, err := gorm.Open(postgres.Open(makePostgresString(cfg)), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Call Migrate function to auto-migrate database schemas
	if cfg.DBMigrate {
		err = Migrate(db)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	// Call See function to auto-migrate database schemas
	if cfg.DBSeeding {
		err = SeedAll(db)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

// makePostgresString creates the PostgreSQL DSN (Data Source Name)
func makePostgresString(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
}
