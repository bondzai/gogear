package rest

import "math"

func PaginateResponse(data []interface{}, count int64, params QueryParams) PaginatedResponse {
	response := PaginatedResponse{
		Items:      data,
		TotalItems: count,
	}

	var totalPages int

	if params.ItemPerPage == 0 {
		totalPages = 1
		params.ItemPerPage = int(count)
	} else {
		totalPages = int(math.Ceil(float64(count) / float64(params.ItemPerPage)))
	}

	response.CurrentPage = params.CurrentPage
	response.ItemsPerPage = params.ItemPerPage
	response.TotalPages = totalPages

	if response.Items == nil {
		response.Items = []interface{}{}
	}

	return response
}
