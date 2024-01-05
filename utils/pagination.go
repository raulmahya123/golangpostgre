package utils

func Pagination(count int64, pageNum, pageSize int) (int, int) {
	totalPages := int(count) / pageSize

	if int(count)%pageSize != 0 {
		totalPages++
	}

	pageSizeNow := pageSize
	if pageSizeNow > int(count) {
		pageSizeNow = int(count)
	}

	return totalPages, pageSizeNow
}
