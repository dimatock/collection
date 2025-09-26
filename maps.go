package collection

// Keys returns a slice containing all the keys of the given map.
// The order of the keys in the returned slice is not guaranteed.
func Keys[K comparable, V any](source map[K]V) []K {
	if source == nil {
		return nil
	}
	result := make([]K, 0, len(source))
	for k := range source {
		result = append(result, k)
	}
	return result
}

// Values returns a slice containing all the values of the given map.
// The order of the values in the returned slice is not guaranteed.
func Values[K comparable, V any](source map[K]V) []V {
	if source == nil {
		return nil
	}
	result := make([]V, 0, len(source))
	for _, v := range source {
		result = append(result, v)
	}
	return result
}

// MapValues creates a new map with the same keys as the source map, but with values transformed
// by the given transform function.
func MapValues[K comparable, V, U any](source map[K]V, transform func(val V) U) map[K]U {
	if source == nil {
		return nil
	}
	result := make(map[K]U, len(source))
	for k, v := range source {
		result[k] = transform(v)
	}
	return result
}
