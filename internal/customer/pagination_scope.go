package customer

import "gorm.io/gorm"

// Scope to paginate in domain
// Doc: https://gorm.io/docs/scopes.html
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Don't accept negative values
		if page <= 0 {
			page = 1
		}

		// Hard limit for pageSize. PageSize less than 100 and negative value protection
		switch {
		case pageSize > 100:
			pageSize = 100

		case page <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.
			Offset(offset).
			Limit(pageSize)
	}
}

// A closure is a function that has another func inside, this is usefull because
// you can pass dynamic variables
