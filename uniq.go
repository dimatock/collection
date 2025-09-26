package collection

// Uniq returns a new slice with duplicate values removed.
// The order of elements is preserved from the original slice.
// It uses a map to track seen elements, so the type T must be comparable.
func Uniq[T comparable](source []T) []T {
	if source == nil {
		return nil
	}

	seen := make(map[T]struct{}, len(source))
	result := make([]T, 0, len(source))

	for _, item := range source {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
