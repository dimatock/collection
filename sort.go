package collection

import "sort"

// Sort returns a new sorted slice based on the provided less function.
// It does not modify the original slice.
func Sort[T any](source []T, less func(a, b T) bool) []T {
	if source == nil {
		return nil
	}

	// Create a copy to preserve the original slice (immutability)
	result := make([]T, len(source))
	copy(result, source)

	// sort.Slice works with a boolean less function, which matches our API.
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	return result
}
