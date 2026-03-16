package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rawndawn/customer-notification/internal/database"
	"github.com/rawndawn/customer-notification/internal/repositories"
	"github.com/rawndawn/customer-notification/internal/scheduler"
	"github.com/rawndawn/customer-notification/internal/services"
)

func main() {
	// Lgger instance to pass for DI
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Starting promotional email sender")

	// Load env
	err := godotenv.Load()

	if err != nil {
		logger.Error("Cannot read .env")
		os.Exit(1)
	}

	// Instance db, repository and all logic for customer promotional email
	database := database.NewDBConnection(logger)
	repository := repositories.NewCustomerRepository(database.ConnectMySQL())
	service := services.NewCustomerService(repository, logger)

	// Scheduler
	monthlyScheduler := scheduler.NewPromotionalEmailScheduler(
		12, // hour (24 format)
		21, // minutes
		service.ProcessMontlyPromotionalEmail,
	)

	// Init scheduler
	monthlyScheduler.Start()

	logger.Error("Scheduler stopped")
}
