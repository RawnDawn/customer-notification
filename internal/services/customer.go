package services

import (
	"log/slog"

	"github.com/rawndawn/customer-notification/internal/models"
	"github.com/rawndawn/customer-notification/internal/repositories"
)

type CustomerService struct {
	repository *repositories.CustomerRepository
	logger     *slog.Logger
}

func NewCustomerService(repository *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: repository,
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
