package repositories

import (
	"github.com/rawndawn/customer-notification/internal/models"
	"github.com/rawndawn/customer-notification/internal/database"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		DB: db,
	}
}

// Get customers using paginate scope
func (r *CustomerRepository) PaginateCustomers(page, pageSize int) ([]models.Customer, error) {
	var customers []models.Customer

	err := r.DB.
		Scopes(database.Paginate(page, pageSize)).
		Find(&customers).
		Error

	return customers, err
}

// TODO - Define the create customer function
