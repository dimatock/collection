package collection

// Some returns true if at least one element in the slice satisfies the predicate.
// It stops iterating and returns true as soon as a match is found.
func Some[T any](source []T, predicate func(item T) bool) bool {
	for _, item := range source {
		if predicate(item) {
			return true
		}
	}
	return false
}
