package collection

// Filter returns a new slice containing only the elements from the source slice
// for which the predicate function returns true.
func Filter[T any](source []T, predicate func(item T) bool) []T {
	if source == nil {
		return nil
	}
	result := make([]T, 0)
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
