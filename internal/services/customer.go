package services

import (
	"log/slog"
	"net/mail"

	"github.com/rawndawn/customer-notification/internal/models"
	"github.com/rawndawn/customer-notification/internal/repositories"
	"github.com/rawndawn/customer-notification/internal/workers"
)

type CustomerService struct {
	repository *repositories.CustomerRepository
	logger     *slog.Logger
}

func NewCustomerService(
	repository *repositories.CustomerRepository,
	logger *slog.Logger,
) *CustomerService {
	return &CustomerService{
		repository: repository,
		logger:     logger,
	}
}

// Get Customers with pagination
func (s *CustomerService) PaginateCustomer(page, pageSize int) ([]models.Customer, error) {
	customers, err := s.repository.PaginateCustomers(page, pageSize)

	if err != nil {
		s.logger.Error("Cannot paginate in Customer Service", slog.Any("err", err))
	}

	return customers, nil
}

// Method to send monthly promotional email
// This iterate in customer pagination to send the workers
func (s *CustomerService) ProcessMontlyPromotionalEmail() {
	// Handling pagination
	var pageSize int = 100

	// Iterate in all customers
	customerCount, err := s.repository.CountCustomers()
	totalPages := int((customerCount + int64(pageSize) - 1) / int64(pageSize))

	if err != nil {
		s.logger.Error("Cannot get customer count in Customer Service", slog.Any("err", err))
	}

	s.logger.Info("Initializing customer monthly promotional email")

	// Here, it's good send workers per page, because internally we use a
	// wait group, so, we don't have memory pressure
	for page := 1; page <= totalPages; page++ {
		// Get customers using pagination
		customers, err := s.PaginateCustomer(page, pageSize)
		if err != nil {
			s.logger.Error(
				"Cannot iterate to send promotional email in Customer Service",
				slog.Any("err", err),
			)
		}

		// Send workers
		workers.StartPromotionalEmailWorkerPool(
			5,
			customers,
		)
	}
}

// Method to validate if a string is a valid email
// Returns true when is a valid email
func (s *CustomerService) IsValidEmail(email string) bool { 
	_, err := mail.ParseAddress(email)

	return err == nil
}