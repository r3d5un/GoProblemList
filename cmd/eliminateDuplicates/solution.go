package elimindateduplicates

func ElimindateDuplicates[T comparable](values []T) []T {
	exists := make(map[T]bool)
	deduplicated := []T{}

	for _, value := range values {
		if _, ok := exists[value]; !ok {
			exists[value] = true
			deduplicated = append(deduplicated, value)
		}
	}

	return deduplicated
}
