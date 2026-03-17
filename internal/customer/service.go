package customer

import (
	"log/slog"
	"net/mail"

	"github.com/rawndawn/customer-notification/internal/notification"
	"github.com/rawndawn/customer-notification/internal/infrestructure/worker"
)

type CustomerService struct {
	repository *CustomerRepository
	logger     *slog.Logger
}

func NewCustomerService(
	repository *CustomerRepository,
	logger     *slog.Logger,
) *CustomerService {
	return &CustomerService{
		repository: repository,
		logger:     logger,
	}
}

// Get Customers with pagination
func (s *CustomerService) PaginateCustomerWithEmail(page, pageSize int) ([]Customer, error) {
	// paginate customers with valid email
	customers, err := s.repository.QueryCustomers(
		WithEmailNotNull,
		Paginate(page, pageSize),
	)

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
		customers, err := s.PaginateCustomerWithEmail(page, pageSize)
		if err != nil {
			s.logger.Error(
				"Cannot iterate to send promotional email in Customer Service",
				slog.Any("err", err),
			)
		}

		worker.StartWorkerPool(
			5,
			customers,
			s.SendPromotionalEmail,
		)
	}
}

// Service method to validate mail and send promotional email to the customer
func (s *CustomerService) SendPromotionalEmail(customer Customer) error {
	// A this point, customer must be has email, that's why we only use return
	if customer.Email == nil {
		return ErrInvalidCustomerEmail
	}

	customerEmail, err := mail.ParseAddress(*customer.Email)

	if err != nil {
		return ErrInvalidCustomerEmail
	}

	err = notification.SendPromotional(customerEmail, customer.Firstname)

	if err != nil {
		s.logger.Error("Cannot send email", slog.Any("err", err))
		return ErrCannotSendEmail
	}

	return nil
}
