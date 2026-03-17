package customer

import (
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
) ([]Customer, error) {
	var customers []Customer

	// From the scopes (filters) add to the query
	query := r.DB.Model(&Customer{}).
		Scopes(scopes...)

	err := query.Find(&customers).Error

	return customers, err
}

// Get total count of customers records
func (r *CustomerRepository) CountCustomers() (int64, error) {
	var count int64

	err := r.DB.Model(&Customer{}).
		Count(&count).
		Error

	return count, err
}
