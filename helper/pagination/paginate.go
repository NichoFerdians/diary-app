package pagination

import (
	"gorm.io/gorm"
)

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		// order := c.DefaultQuery("order_by", "id asc")

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize) //.Order(order)
	}
}
