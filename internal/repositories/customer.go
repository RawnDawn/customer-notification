package repositories

import (
	"github.com/rawndawn/customer-notification/internal/models"
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

// Get customers passing scopes to the query
func (r *CustomerRepository) QueryCustomers(
	scopes ...func(*gorm.DB) *gorm.DB,
) ([]models.Customer, error) {
	var customers []models.Customer

	// From the scopes (filters) add to the query
	query := r.DB.Model(&models.Customer{}).
		Scopes(scopes...)

	err := query.Find(&customers).Error

	return customers, err
}

// Get total count of customers records
func (r *CustomerRepository) CountCustomers() (int64, error) {
	var count int64

	err := r.DB.Model(&models.Customer{}).
		Count(&count).
		Error

	return count, err
}
