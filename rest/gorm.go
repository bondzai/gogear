package rest

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ApplySorting adds sorting to the query based on QueryParams
func ApplySorting(db *gorm.DB, sortKey, sortPattern string) *gorm.DB {
	if sortKey != "" {
		// Apply sort pattern, defaults to ascending if not specified
		sortOrder := "ASC"
		if strings.ToLower(sortPattern) == "desc" {
			sortOrder = "DESC"
		}
		return db.Order(fmt.Sprintf("%s %s", sortKey, sortOrder))
	}

	return db
}

// ApplyPagination applies pagination logic based on QueryParams
func ApplyPagination(db *gorm.DB, currentPage, itemPerPage int) *gorm.DB {
	if itemPerPage <= 0 {
		itemPerPage = 10 // Default items per page
	}

	offset := (currentPage - 1) * itemPerPage
	return db.Offset(offset).Limit(itemPerPage)
}

// ApplyFilters adds filtering logic to the query based on QueryParams
func ApplyFilters(db *gorm.DB, filterParams map[string]interface{}, searchParams map[string]string) *gorm.DB {
	// Apply normal filters
	for key, value := range filterParams {
		if s, ok := value.(string); ok && s == "null" {
			db = db.Where(fmt.Sprintf("%s is null", key))
		} else {
			db = db.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	// Apply search filters using OR logic
	if len(searchParams) > 0 {
		var orConditions []string
		var orValues []interface{}

		// Build OR conditions for search parameters
		for column, searchValue := range searchParams {
			orConditions = append(orConditions, fmt.Sprintf("%s ILIKE ?", column))
			orValues = append(orValues, "%"+searchValue+"%")
		}

		// Combine the conditions into a single query using OR
		db = db.Where("("+strings.Join(orConditions, " OR ")+")", orValues...)
	}

	return db
}
