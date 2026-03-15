package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type DBConnection struct {
	logger *slog.Logger
}

func NewDBConnection(logger *slog.Logger) *DBConnection {
	return &DBConnection{
		logger: logger,
	}
}

// Connect to MySQL
func (con *DBConnection) ConnectMySQL() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		con.logger.Error("Cannot connect to MySQL", slog.Any("err", err))
		os.Exit(1)
	}

	con.logger.Info("Connect to MySQL")

	return db
}
