package utils

import (
	"fmt"
	"math"
)

// Paginate returns a slice of items for a given page number
func Paginate(items []interface{}, pageNumber, pageSize int) ([]interface{}, error) {
	totalItems := len(items)
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	if pageNumber < 1 || pageNumber > totalPages {
		return nil, fmt.Errorf("Invalid page number")
	}

	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > totalItems {
		endIndex = totalItems
	}

	return items[startIndex:endIndex], nil
}
