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

// Get customers using paginate scope
func (r *CustomerRepository) PaginateCustomers(page, pageSize int) ([]models.Customer, error) {
	var customers []models.Customer

	err := r.DB.
		Scopes(Paginate(page, pageSize)).
		Find(&customers).
		Error

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
