package models

import "time"

type Customer struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Firstname string     `json:"first_name" gorm:"column:first_name;size:120;not null"`
	Lastname  string     `json:"last_name" gorm:"column:last_name;size:120;not null"`
	Phone     string     `json:"phone" gorm:"size:20"`
	Email     string     `json:"email" gorm:"not null;uniqueIndex;size:255"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

// Specification for the real database table
func (Customer) TableName() string {
	return "erp_webservice_customers"
}
