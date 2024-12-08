package utils

// AllPairs generates all unique pairs from a slice of any type.
func AllPairs[T any](input []T) [][]T {
	var pairs [][]T

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
		    if i != j {
			    pairs = append(pairs, []T{input[i], input[j]})
		    }
		}
	}

	return pairs
}
