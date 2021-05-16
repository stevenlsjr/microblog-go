package api

type PaginationParams struct {
	Limit    int64   `json:"limit"`
	Offset   int64   `json:"offset"`
	Count    int64   `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}
