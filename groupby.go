package collection

// GroupBy groups elements from the source slice into a map.
// The key of the map is determined by the keySelector function, and the value is a slice of elements
// that share the same key. The key type K must be comparable.
func GroupBy[T any, K comparable](source []T, keySelector func(item T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range source {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}
