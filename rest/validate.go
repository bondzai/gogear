package rest

import (
	"strings"
)

func Validate(params map[string]string) (*QueryParams, error) {
	var page, pageSize *int

	// Sorting
	sortParams := params[SortKey]
	if sortParams != "" {
		sortFields := strings.Split(sortParams, ",")
		for i, sortField := range sortFields {
			sortFields[i] = strings.ToLower(sortField)
			if sortField == "ID" {
				sortFields[i] = "_id"
			}
		}
		sortParams = strings.Join(sortFields, ",")
	}

	sortPatternParams := params[SortPatternKey]

	// filter
	filterParams := make(map[string]interface{})
	for key, value := range params {
		if key == SortKey || key == SortPatternKey || key == CurrentPageKey || key == ItemsPerPageKey {
			continue
		}

		if isInSlice(SpecialKeys, key) {
			filterParams[camelToSnakeCase(key)] = value
			continue
		}

		values := strings.Split(value, ",")
		if len(values) > 1 {
			filterParams[camelToSnakeCase(key)] = values
		} else {
			filterParams[camelToSnakeCase(key)] = value
		}
	}

	defaultPage := 1
	page = &defaultPage

	defaultSize := 10
	pageSize = &defaultSize

	return &QueryParams{
		FilterParams:      filterParams,
		SortParams:        sortParams,
		SortPatternParams: sortPatternParams,
		CurrentPage:       page,
		ItemPerPage:       pageSize,
	}, nil
}
