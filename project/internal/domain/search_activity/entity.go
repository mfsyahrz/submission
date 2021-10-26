package search_activity

import "time"

type SearchActivity struct {
	ID        int64
	Activity  string
	Success   bool
	Message   string
	JsonData  string
	RequestDt time.Time
}
