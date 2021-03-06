package models

import (
	"net/url"
	"strconv"
)

// OrderDirection represents ascending or descending direction when sorting
type OrderDirection byte

const (
	ASCENDING OrderDirection = iota
	DESCENDING
)

// ListPageRequest contains filters for querying lists
type ListPageRequest struct {
	Page           int
	PageSize       int
	OrderBy        string
	OrderDirection OrderDirection
	SearchQuery    string
	SearchInverted bool
	Extension      string
}

// FromQueryString initializes page request from query sting values
func (p *ListPageRequest) FromQueryString(q *url.Values, availableColumns []string) {

	if page, err := strconv.Atoi(q.Get("Page")); err == nil {
		p.Page = page
	}

	if pageSize, err := strconv.Atoi(q.Get("PageSize")); err == nil {
		p.PageSize = pageSize
	}

	//protect OrderBy from sql injection
	orderByParam := q.Get("OrderBy")
	for _, col := range availableColumns {
		if col == orderByParam {
			p.OrderBy = col
		}
	}

	if dir, err := strconv.Atoi(q.Get("OrderDirection")); err == nil {
		p.OrderDirection = OrderDirection(dir)
	}

	p.SearchQuery = q.Get("SearchQuery")

	if searchInverted, err := strconv.ParseBool(q.Get("SearchInverted")); err == nil {
		p.SearchInverted = searchInverted
	}

	p.Extension = q.Get("Extension")
}

// IsValid validates input from the user
func (p *ListPageRequest) IsValid() bool {
	if len(p.SearchQuery) > 0 && len(p.SearchQuery) < 3 {
		return false
	}

	return true
}
