package util

func SetPageForPagination(page int64) int64 {
	if page == 0 {
		page = 1
	}
	return page
}
