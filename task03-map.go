package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	values := make([]string, 0, len(input))

	for i := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, key := range keys {
		values = append(values, input[key])
	}

	return values
}
