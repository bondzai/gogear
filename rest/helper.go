package rest

import (
	"unicode"
)

func camelToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func isInSlice[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}
