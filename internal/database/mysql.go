package database

import (
	"os"
	"log/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection struct {
	logger *slog.Logger
}

// Connect to MySQL
func (con *DBConnection) ConnectMySQL() *gorm.DB { 
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil { 
		con.logger.Error("Cannot connect to MySQL", slog.Any("err", err))
	}

	con.logger.Info("Connect to MySQL")

	return db
}