package rest

import (
	"strconv"
	"strings"
)

const (
	DefaultDateFormat = "2006-01-02"
	CreatedAtKey      = "createdAt"
	SortKey           = "sortBy"
	SortPatternKey    = "pattern"
	CurrentPageKey    = "page"
	ItemsPerPageKey   = "size"
	DefaultPage       = 1
	DefaultSize       = 10
)

var (
	SearchWithRegexKey = []string{"search"}
	ReserveKeys        = []string{
		SortKey,
		SortPatternKey,
		CurrentPageKey,
		ItemsPerPageKey,
	}
)

func BuildQueryParams(params map[string]string, searchColumns ...string) *QueryParams {
	// sort
	sortParams := params[SortKey]
	if sortParams != "" {
		sortFields := strings.Split(sortParams, ",")

		for i, sortField := range sortFields {
			sortFields[i] = strings.ToLower(sortField)
		}

		sortParams = strings.Join(sortFields, ",")
	}
	sortPatternParams := params[SortPatternKey]

	// filter
	filterParams := make(map[string]interface{})
	searchParams := make(map[string]string)

	for key, value := range params {
		if isInSlice(ReserveKeys, key) {
			continue
		}

		if isInSlice(SearchWithRegexKey, key) {
			for _, col := range searchColumns {
				searchParams[col] = value
			}
			continue
		}

		values := strings.Split(value, ",")
		if len(values) > 1 {
			filterParams[camelToSnakeCase(key)] = values
		} else {
			filterParams[camelToSnakeCase(key)] = value
		}
	}

	parseIntWithDefault := func(key string, defaultValue int) int {
		if val, exists := params[key]; exists {
			if num, err := strconv.Atoi(val); err == nil {
				return num
			}
		}
		return defaultValue
	}

	page := parseIntWithDefault(CurrentPageKey, DefaultPage)
	pageSize := parseIntWithDefault(ItemsPerPageKey, DefaultSize)

	return &QueryParams{
		FilterParams:      filterParams,
		SearchParams:      searchParams,
		SortParams:        sortParams,
		SortPatternParams: sortPatternParams,
		CurrentPage:       page,
		ItemPerPage:       pageSize,
	}
}
