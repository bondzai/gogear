package rest

type QueryParams struct {
	FilterParams      map[string]interface{}
	SortParams        string `json:"sortBy,omitempty"`
	SortPatternParams string `json:"pattern,omitempty"`
	CurrentPage       *int   `json:"page,omitempty"`
	ItemPerPage       *int   `json:"size,omitempty"`
}

type PaginatedResponse struct {
	Items        interface{} `json:"items"`
	TotalItems   int64       `json:"total"`
	TotalPages   int         `json:"totalPage"`
	CurrentPage  *int        `json:"page,omitempty"`
	ItemsPerPage *int        `json:"size,omitempty"`
}

const (
	DefaultDateFormat = "2006-01-02"
	CreatedAtKey      = "createdAt"
	SortKey           = "sortBy"
	SortPatternKey    = "pattern"
	CurrentPageKey    = "page"
	ItemsPerPageKey   = "size"
)

// Filter with Regex
var SpecialKeys = []string{}
