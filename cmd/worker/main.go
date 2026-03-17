package main

import (
	"log/slog"
	"os"
	"strconv"

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
	hour, err := strconv.Atoi(os.Getenv("SCHEDULER_TIME_HOUR"))
	if err != nil {
		logger.Error("Cannot parse SCHEDULER_TIME_HOUR into int")
		os.Exit(1)
	}

	minute, err := strconv.Atoi(os.Getenv("SCHEDULER_TIME_MINUTE"))
	if err != nil {
		logger.Error("Cannot parse SCHEDULER_TIME_MINUTE into int")
		os.Exit(1)
	}

	monthlyScheduler := scheduler.NewPromotionalEmailScheduler(
		hour, // hour (24 format)
		minute, // minutes
		service.ProcessMontlyPromotionalEmail,
	)

	// Init scheduler
	monthlyScheduler.Start()

	logger.Error("Scheduler stopped")
}
