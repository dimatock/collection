package collection

// Every returns true if all elements in the slice satisfy the predicate.
// It stops iterating and returns false as soon as a mismatch is found.
func Every[T any](source []T, predicate func(item T) bool) bool {
	for _, item := range source {
		if !predicate(item) {
			return false
		}
	}
	return true
}
