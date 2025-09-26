package collection

// Find returns the first element in the slice that satisfies the predicate.
// It returns the element and true if found, otherwise it returns the zero value of the element type and false.
func Find[T any](source []T, predicate func(item T) bool) (T, bool) {
	for _, item := range source {
		if predicate(item) {
			return item, true
		}
	}

	var zero T
	return zero, false
}
