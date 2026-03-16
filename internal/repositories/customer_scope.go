package repositories

import "gorm.io/gorm"

func WithEmailNotNull(db *gorm.DB) *gorm.DB {
	return db.
		Where("email IS NOT NULL").
		Where("email <> ''")
}
