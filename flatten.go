package collection

// Flatten converts a slice of slices into a single slice.
func Flatten[T any](source [][]T) []T {
	var totalLen int
	for _, s := range source {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)
	for _, s := range source {
		result = append(result, s...)
	}
	return result
}
